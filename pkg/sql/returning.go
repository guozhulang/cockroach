// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql

import (
	"bytes"

	"golang.org/x/net/context"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/pkg/errors"
)

// returningHelper implements the logic used for statements with RETURNING clauses. It accumulates
// result rows, one for each call to append().
type returningHelper struct {
	p *planner
	// Expected columns.
	columns sqlbase.ResultColumns
	// Processed copies of expressions from ReturningExprs.
	exprs        []tree.TypedExpr
	rowCount     int
	source       *dataSourceInfo
	curSourceRow tree.Datums

	// This struct must be allocated on the heap and its location stay
	// stable after construction because it implements
	// IndexedVarContainer and the IndexedVar objects in sub-expressions
	// will link to it by reference after checkRenderStar / analyzeExpr.
	// Enforce this using NoCopy.
	//
	//lint:ignore U1000 this marker prevents by-value copies.
	noCopy util.NoCopy
}

// newReturningHelper creates a new returningHelper for use by an
// insert/update node.
func (p *planner) newReturningHelper(
	ctx context.Context,
	r tree.ReturningClause,
	desiredTypes []types.T,
	tn *tree.TableName,
	tablecols []sqlbase.ColumnDescriptor,
) (*returningHelper, error) {
	rh := &returningHelper{
		p: p,
	}
	var rExprs tree.ReturningExprs
	switch t := r.(type) {
	case *tree.ReturningExprs:
		rExprs = *t
	case *tree.ReturningNothing, *tree.NoReturningClause:
		return rh, nil
	default:
		panic(errors.Errorf("unexpected ReturningClause type: %T", t))
	}

	for _, e := range rExprs {
		if err := p.txCtx.AssertNoAggregationOrWindowing(
			e.Expr, "RETURNING", p.session.SearchPath,
		); err != nil {
			return nil, err
		}
	}

	rh.columns = make(sqlbase.ResultColumns, 0, len(rExprs))
	rh.source = newSourceInfoForSingleTable(
		*tn, sqlbase.ResultColumnsFromColDescs(tablecols),
	)
	rh.exprs = make([]tree.TypedExpr, 0, len(rExprs))
	ivarHelper := tree.MakeIndexedVarHelper(rh, len(tablecols))
	for _, target := range rExprs {
		cols, typedExprs, _, err := p.computeRenderAllowingStars(
			ctx, target, types.Any, multiSourceInfo{rh.source}, ivarHelper,
			autoGenerateRenderOutputName)
		if err != nil {
			return nil, err
		}
		rh.columns = append(rh.columns, cols...)
		rh.exprs = append(rh.exprs, typedExprs...)
	}
	return rh, nil
}

// cookResultRow prepares a row according to the ReturningExprs, with input values
// from rowVals.
func (rh *returningHelper) cookResultRow(rowVals tree.Datums) (tree.Datums, error) {
	if rh.exprs == nil {
		rh.rowCount++
		return rowVals, nil
	}
	rh.curSourceRow = rowVals
	resRow := make(tree.Datums, len(rh.exprs))
	for i, e := range rh.exprs {
		d, err := e.Eval(&rh.p.evalCtx)
		if err != nil {
			return nil, err
		}
		resRow[i] = d
	}
	return resRow, nil
}

// IndexedVarEval implements the tree.IndexedVarContainer interface.
func (rh *returningHelper) IndexedVarEval(idx int, ctx *tree.EvalContext) (tree.Datum, error) {
	return rh.curSourceRow[idx].Eval(ctx)
}

// IndexedVarResolvedType implements the tree.IndexedVarContainer interface.
func (rh *returningHelper) IndexedVarResolvedType(idx int) types.T {
	return rh.source.sourceColumns[idx].Typ
}

// IndexedVarFormat implements the tree.IndexedVarContainer interface.
func (rh *returningHelper) IndexedVarFormat(buf *bytes.Buffer, f tree.FmtFlags, idx int) {
	rh.source.FormatVar(buf, f, idx)
}
