// Code generated by ent, DO NOT EDIT.

package subtask

import (
	"sonar-bat/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldCreatedBy, v))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldTaskID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldStatus, v))
}

// DetectionType applies equality check predicate on the "detection_type" field. It's identical to DetectionTypeEQ.
func DetectionType(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldDetectionType, v))
}

// SrcEpSelNum applies equality check predicate on the "src_ep_sel_num" field. It's identical to SrcEpSelNumEQ.
func SrcEpSelNum(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldSrcEpSelNum, v))
}

// DstEpSelNum applies equality check predicate on the "dst_ep_sel_num" field. It's identical to DstEpSelNumEQ.
func DstEpSelNum(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldDstEpSelNum, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldCreatedBy, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...uuid.UUID) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldTaskID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldStatus, v))
}

// DetectionTypeEQ applies the EQ predicate on the "detection_type" field.
func DetectionTypeEQ(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldDetectionType, v))
}

// DetectionTypeNEQ applies the NEQ predicate on the "detection_type" field.
func DetectionTypeNEQ(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldDetectionType, v))
}

// DetectionTypeIn applies the In predicate on the "detection_type" field.
func DetectionTypeIn(vs ...int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldDetectionType, vs...))
}

// DetectionTypeNotIn applies the NotIn predicate on the "detection_type" field.
func DetectionTypeNotIn(vs ...int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldDetectionType, vs...))
}

// DetectionTypeGT applies the GT predicate on the "detection_type" field.
func DetectionTypeGT(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldDetectionType, v))
}

// DetectionTypeGTE applies the GTE predicate on the "detection_type" field.
func DetectionTypeGTE(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldDetectionType, v))
}

// DetectionTypeLT applies the LT predicate on the "detection_type" field.
func DetectionTypeLT(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldDetectionType, v))
}

// DetectionTypeLTE applies the LTE predicate on the "detection_type" field.
func DetectionTypeLTE(v int16) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldDetectionType, v))
}

// SrcEpSelNumEQ applies the EQ predicate on the "src_ep_sel_num" field.
func SrcEpSelNumEQ(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldSrcEpSelNum, v))
}

// SrcEpSelNumNEQ applies the NEQ predicate on the "src_ep_sel_num" field.
func SrcEpSelNumNEQ(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldSrcEpSelNum, v))
}

// SrcEpSelNumIn applies the In predicate on the "src_ep_sel_num" field.
func SrcEpSelNumIn(vs ...int) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldSrcEpSelNum, vs...))
}

// SrcEpSelNumNotIn applies the NotIn predicate on the "src_ep_sel_num" field.
func SrcEpSelNumNotIn(vs ...int) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldSrcEpSelNum, vs...))
}

// SrcEpSelNumGT applies the GT predicate on the "src_ep_sel_num" field.
func SrcEpSelNumGT(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldSrcEpSelNum, v))
}

// SrcEpSelNumGTE applies the GTE predicate on the "src_ep_sel_num" field.
func SrcEpSelNumGTE(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldSrcEpSelNum, v))
}

// SrcEpSelNumLT applies the LT predicate on the "src_ep_sel_num" field.
func SrcEpSelNumLT(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldSrcEpSelNum, v))
}

// SrcEpSelNumLTE applies the LTE predicate on the "src_ep_sel_num" field.
func SrcEpSelNumLTE(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldSrcEpSelNum, v))
}

// DstEpSelNumEQ applies the EQ predicate on the "dst_ep_sel_num" field.
func DstEpSelNumEQ(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldEQ(FieldDstEpSelNum, v))
}

// DstEpSelNumNEQ applies the NEQ predicate on the "dst_ep_sel_num" field.
func DstEpSelNumNEQ(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldNEQ(FieldDstEpSelNum, v))
}

// DstEpSelNumIn applies the In predicate on the "dst_ep_sel_num" field.
func DstEpSelNumIn(vs ...int) predicate.Subtask {
	return predicate.Subtask(sql.FieldIn(FieldDstEpSelNum, vs...))
}

// DstEpSelNumNotIn applies the NotIn predicate on the "dst_ep_sel_num" field.
func DstEpSelNumNotIn(vs ...int) predicate.Subtask {
	return predicate.Subtask(sql.FieldNotIn(FieldDstEpSelNum, vs...))
}

// DstEpSelNumGT applies the GT predicate on the "dst_ep_sel_num" field.
func DstEpSelNumGT(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldGT(FieldDstEpSelNum, v))
}

// DstEpSelNumGTE applies the GTE predicate on the "dst_ep_sel_num" field.
func DstEpSelNumGTE(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldGTE(FieldDstEpSelNum, v))
}

// DstEpSelNumLT applies the LT predicate on the "dst_ep_sel_num" field.
func DstEpSelNumLT(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldLT(FieldDstEpSelNum, v))
}

// DstEpSelNumLTE applies the LTE predicate on the "dst_ep_sel_num" field.
func DstEpSelNumLTE(v int) predicate.Subtask {
	return predicate.Subtask(sql.FieldLTE(FieldDstEpSelNum, v))
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.Subtask {
	return predicate.Subtask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.Subtask {
	return predicate.Subtask(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subtask) predicate.Subtask {
	return predicate.Subtask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subtask) predicate.Subtask {
	return predicate.Subtask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subtask) predicate.Subtask {
	return predicate.Subtask(sql.NotPredicates(p))
}
