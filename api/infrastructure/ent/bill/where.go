// Code generated by ent, DO NOT EDIT.

package bill

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldID, id))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldPrice, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Bill {
	return predicate.Bill(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Bill {
	return predicate.Bill(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Bill {
	return predicate.Bill(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Bill {
	return predicate.Bill(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Bill {
	return predicate.Bill(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Bill {
	return predicate.Bill(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Bill {
	return predicate.Bill(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Bill {
	return predicate.Bill(sql.FieldLTE(FieldPrice, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSrcUser applies the HasEdge predicate on the "src_user" edge.
func HasSrcUser() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, SrcUserTable, SrcUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSrcUserWith applies the HasEdge predicate on the "src_user" edge with a given conditions (other predicates).
func HasSrcUserWith(preds ...predicate.User) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newSrcUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDstUser applies the HasEdge predicate on the "dst_user" edge.
func HasDstUser() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DstUserTable, DstUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDstUserWith applies the HasEdge predicate on the "dst_user" edge with a given conditions (other predicates).
func HasDstUserWith(preds ...predicate.User) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := newDstUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bill) predicate.Bill {
	return predicate.Bill(sql.NotPredicates(p))
}
