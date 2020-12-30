// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/B6001186/Contagions/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AreaQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AreaQueryRuleFunc func(context.Context, *ent.AreaQuery) error

// EvalQuery return f(ctx, q).
func (f AreaQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AreaQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AreaQuery", q)
}

// The AreaMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AreaMutationRuleFunc func(context.Context, *ent.AreaMutation) error

// EvalMutation calls f(ctx, m).
func (f AreaMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AreaMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AreaMutation", m)
}

// The BloodtypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BloodtypeQueryRuleFunc func(context.Context, *ent.BloodtypeQuery) error

// EvalQuery return f(ctx, q).
func (f BloodtypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BloodtypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BloodtypeQuery", q)
}

// The BloodtypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BloodtypeMutationRuleFunc func(context.Context, *ent.BloodtypeMutation) error

// EvalMutation calls f(ctx, m).
func (f BloodtypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BloodtypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BloodtypeMutation", m)
}

// The CategoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CategoryQueryRuleFunc func(context.Context, *ent.CategoryQuery) error

// EvalQuery return f(ctx, q).
func (f CategoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CategoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CategoryQuery", q)
}

// The CategoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CategoryMutationRuleFunc func(context.Context, *ent.CategoryMutation) error

// EvalMutation calls f(ctx, m).
func (f CategoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CategoryMutation", m)
}

// The DepartmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DepartmentQueryRuleFunc func(context.Context, *ent.DepartmentQuery) error

// EvalQuery return f(ctx, q).
func (f DepartmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DepartmentQuery", q)
}

// The DepartmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DepartmentMutationRuleFunc func(context.Context, *ent.DepartmentMutation) error

// EvalMutation calls f(ctx, m).
func (f DepartmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DepartmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DepartmentMutation", m)
}

// The DiagnosisQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DiagnosisQueryRuleFunc func(context.Context, *ent.DiagnosisQuery) error

// EvalQuery return f(ctx, q).
func (f DiagnosisQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DiagnosisQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DiagnosisQuery", q)
}

// The DiagnosisMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DiagnosisMutationRuleFunc func(context.Context, *ent.DiagnosisMutation) error

// EvalMutation calls f(ctx, m).
func (f DiagnosisMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DiagnosisMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DiagnosisMutation", m)
}

// The DiseaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DiseaseQueryRuleFunc func(context.Context, *ent.DiseaseQuery) error

// EvalQuery return f(ctx, q).
func (f DiseaseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DiseaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DiseaseQuery", q)
}

// The DiseaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DiseaseMutationRuleFunc func(context.Context, *ent.DiseaseMutation) error

// EvalMutation calls f(ctx, m).
func (f DiseaseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DiseaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DiseaseMutation", m)
}

// The DiseasetypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DiseasetypeQueryRuleFunc func(context.Context, *ent.DiseasetypeQuery) error

// EvalQuery return f(ctx, q).
func (f DiseasetypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DiseasetypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DiseasetypeQuery", q)
}

// The DiseasetypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DiseasetypeMutationRuleFunc func(context.Context, *ent.DiseasetypeMutation) error

// EvalMutation calls f(ctx, m).
func (f DiseasetypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DiseasetypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DiseasetypeMutation", m)
}

// The DrugQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DrugQueryRuleFunc func(context.Context, *ent.DrugQuery) error

// EvalQuery return f(ctx, q).
func (f DrugQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DrugQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DrugQuery", q)
}

// The DrugMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DrugMutationRuleFunc func(context.Context, *ent.DrugMutation) error

// EvalMutation calls f(ctx, m).
func (f DrugMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DrugMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DrugMutation", m)
}

// The DrugTypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DrugTypeQueryRuleFunc func(context.Context, *ent.DrugTypeQuery) error

// EvalQuery return f(ctx, q).
func (f DrugTypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DrugTypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DrugTypeQuery", q)
}

// The DrugTypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DrugTypeMutationRuleFunc func(context.Context, *ent.DrugTypeMutation) error

// EvalMutation calls f(ctx, m).
func (f DrugTypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DrugTypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DrugTypeMutation", m)
}

// The EmployeeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmployeeQueryRuleFunc func(context.Context, *ent.EmployeeQuery) error

// EvalQuery return f(ctx, q).
func (f EmployeeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmployeeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmployeeQuery", q)
}

// The EmployeeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmployeeMutationRuleFunc func(context.Context, *ent.EmployeeMutation) error

// EvalMutation calls f(ctx, m).
func (f EmployeeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmployeeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmployeeMutation", m)
}

// The GenderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenderQueryRuleFunc func(context.Context, *ent.GenderQuery) error

// EvalQuery return f(ctx, q).
func (f GenderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenderQuery", q)
}

// The GenderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenderMutationRuleFunc func(context.Context, *ent.GenderMutation) error

// EvalMutation calls f(ctx, m).
func (f GenderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenderMutation", m)
}

// The LevelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LevelQueryRuleFunc func(context.Context, *ent.LevelQuery) error

// EvalQuery return f(ctx, q).
func (f LevelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LevelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LevelQuery", q)
}

// The LevelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LevelMutationRuleFunc func(context.Context, *ent.LevelMutation) error

// EvalMutation calls f(ctx, m).
func (f LevelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LevelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LevelMutation", m)
}

// The NametitleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NametitleQueryRuleFunc func(context.Context, *ent.NametitleQuery) error

// EvalQuery return f(ctx, q).
func (f NametitleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NametitleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NametitleQuery", q)
}

// The NametitleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NametitleMutationRuleFunc func(context.Context, *ent.NametitleMutation) error

// EvalMutation calls f(ctx, m).
func (f NametitleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NametitleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NametitleMutation", m)
}

// The PatientQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PatientQueryRuleFunc func(context.Context, *ent.PatientQuery) error

// EvalQuery return f(ctx, q).
func (f PatientQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PatientQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PatientQuery", q)
}

// The PatientMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PatientMutationRuleFunc func(context.Context, *ent.PatientMutation) error

// EvalMutation calls f(ctx, m).
func (f PatientMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PatientMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PatientMutation", m)
}

// The PlaceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PlaceQueryRuleFunc func(context.Context, *ent.PlaceQuery) error

// EvalQuery return f(ctx, q).
func (f PlaceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PlaceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PlaceQuery", q)
}

// The PlaceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PlaceMutationRuleFunc func(context.Context, *ent.PlaceMutation) error

// EvalMutation calls f(ctx, m).
func (f PlaceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PlaceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PlaceMutation", m)
}

// The SeverityQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SeverityQueryRuleFunc func(context.Context, *ent.SeverityQuery) error

// EvalQuery return f(ctx, q).
func (f SeverityQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SeverityQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SeverityQuery", q)
}

// The SeverityMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SeverityMutationRuleFunc func(context.Context, *ent.SeverityMutation) error

// EvalMutation calls f(ctx, m).
func (f SeverityMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SeverityMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SeverityMutation", m)
}

// The StatisticQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StatisticQueryRuleFunc func(context.Context, *ent.StatisticQuery) error

// EvalQuery return f(ctx, q).
func (f StatisticQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StatisticQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StatisticQuery", q)
}

// The StatisticMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StatisticMutationRuleFunc func(context.Context, *ent.StatisticMutation) error

// EvalMutation calls f(ctx, m).
func (f StatisticMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StatisticMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StatisticMutation", m)
}