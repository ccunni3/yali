// Yali provides a framework to implement artribute based access control.
package yali

// New creates a new enforcer for the provided policies.
func New(policies Policies) Enforcer {
	panic("//TODO: not implemented")
	return Enforcer{}
}

// Enforcer enforces configured policies by evaulating request attributes.
// The zero value of an enforcer denies all requests.
type Enforcer struct {
}

// Enforce enforces a poilcy given a set of attributes, returning a Decision to either permit or deny the request.
//
// Permitted requests always return a decision value with Decision.Permit equal to true.
// Denied requests always return a decision value with Decision.Permit equal to false.
func (e *Enforcer) Enforce(Request) Decision {
	panic("//TODO: not implemented")
	//
	// find policy for request
	// apply request to policy
	return Decision{}
}

type Request struct {
	Subject     Subject
	Object      Object
	Action      Action
	Environment Environment
}

// Subject is the set of attributes describing who is performing an action on an object.
type Subject struct {
	ID     string
	Email  string
	Emails []string
	Group  string
	Groups []string
	Role   string
	Roles  []string
}

// Object is the set of attributes describing what a subject is acting on.
type Object struct {
	ID      string
	Type    string
	OwnerID string
}

// Action is the set of attributes describing the action the subject wants to perform on an object.
type Action struct {
	Operation string
}

// Environment is the set of attributes describing the environment in which the subject requested to perform an action on an object.
type Environment struct {
	Now      uint64
	ClientID string
	IPV4     string
}

// Decision is the result of evaluating a request against a policy.
type Decision struct {
	Permit bool
}

// Error implements the Go builtin error interface.
func (d Decision) Error() string {
	panic("//TODO: not implemented")
}

type Policies []Policy

type Policy struct {
	Name        string
	Description string
	Matchers    []Predicate
	OrMatchers  bool
	Rules       []Predicate
	OrRules     bool
	parsed      bool
}

func (p *Policy) Match(r Request) bool {
	if !p.parsed {
		return false
	}
	match, _ := reduce(p.Matchers, r, p.OrMatchers)
	return match
}

func (p *Policy) Evaluate(r Request) Decision {
	if !p.parsed {
		return Decision{Permit: false}
	}
	permit, _ := reduce(p.Rules, r, p.OrRules)
	return Decision{Permit: permit}
}

func (p *Policy) String() string {
	return p.Name
}

type Predicate struct {
	RawExpression string
	test          func(Request) bool
}

func (p *Predicate) String() string {
	return p.RawExpression
}

func reduce(predicates []Predicate, request Request, initialValue bool) (result bool, index int) {
	result = !initialValue
	for i, p := range predicates {
		ok := p.test(request)
		if initialValue {
			result = result || ok
		} else {
			result = result && ok
		}
		if !result {
			index = i
			break
		}
	}
	return
}
