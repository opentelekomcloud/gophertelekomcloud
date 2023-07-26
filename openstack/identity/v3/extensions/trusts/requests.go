package trusts

import "github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/tokens"

// AuthOptsExt extends the base Identity v3 tokens AuthOpts with a TrustID.
type AuthOptsExt struct {
	tokens.AuthOptionsBuilder

	// TrustID is the ID of the trust.
	TrustID string `json:"id"`
}

// ToTokenV3CreateMap builds a create request body from the AuthOpts.
func (opts AuthOptsExt) ToTokenV3CreateMap(scope map[string]any) (map[string]any, error) {
	return opts.AuthOptionsBuilder.ToTokenV3CreateMap(scope)
}

// ToTokenV3ScopeMap builds a scope from AuthOpts.
func (opts AuthOptsExt) ToTokenV3ScopeMap() (map[string]any, error) {
	b, err := opts.AuthOptionsBuilder.ToTokenV3ScopeMap()
	if err != nil {
		return nil, err
	}

	if opts.TrustID != "" {
		if b == nil {
			b = make(map[string]any)
		}
		b["OS-TRUST:trust"] = map[string]any{
			"id": opts.TrustID,
		}
	}

	return b, nil
}

func (opts AuthOptsExt) CanReauth() bool {
	return opts.AuthOptionsBuilder.CanReauth()
}
