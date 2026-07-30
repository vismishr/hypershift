package oauth

import (
	"testing"

	. "github.com/onsi/gomega"

	hyperv1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
)

func TestResolveOAuthVerbosity(t *testing.T) {
	logLevel := func(l hyperv1.LogLevel) hyperv1.OAuthServerOperatorSpec {
		return hyperv1.OAuthServerOperatorSpec{
			ComponentLogLevelSpec: hyperv1.ComponentLogLevelSpec{LogLevel: &l},
		}
	}

	tests := []struct {
		name       string
		hcp        *hyperv1.HostedControlPlane
		expected   int
		expectedOk bool
	}{
		{
			name: "When no operatorConfiguration is set it should return false",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{},
			},
			expected:   0,
			expectedOk: false,
		},
		{
			name: "When operatorConfiguration exists but logLevel is nil it should return false",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{
					OperatorConfiguration: &hyperv1.OperatorConfiguration{},
				},
			},
			expected:   0,
			expectedOk: false,
		},
		{
			name: "When oauthServer logLevel is Normal it should return verbosity 2",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{
					OperatorConfiguration: &hyperv1.OperatorConfiguration{
						OAuthServer: logLevel(hyperv1.Normal),
					},
				},
			},
			expected:   2,
			expectedOk: true,
		},
		{
			name: "When oauthServer logLevel is Debug it should return verbosity 4",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{
					OperatorConfiguration: &hyperv1.OperatorConfiguration{
						OAuthServer: logLevel(hyperv1.Debug),
					},
				},
			},
			expected:   4,
			expectedOk: true,
		},
		{
			name: "When oauthServer logLevel is Trace it should return verbosity 6",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{
					OperatorConfiguration: &hyperv1.OperatorConfiguration{
						OAuthServer: logLevel(hyperv1.Trace),
					},
				},
			},
			expected:   6,
			expectedOk: true,
		},
		{
			name: "When oauthServer logLevel is TraceAll it should return verbosity 8",
			hcp: &hyperv1.HostedControlPlane{
				Spec: hyperv1.HostedControlPlaneSpec{
					OperatorConfiguration: &hyperv1.OperatorConfiguration{
						OAuthServer: logLevel(hyperv1.TraceAll),
					},
				},
			},
			expected:   8,
			expectedOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			v, ok := resolveOAuthVerbosity(tt.hcp)
			g.Expect(ok).To(Equal(tt.expectedOk))
			g.Expect(v).To(Equal(tt.expected))
		})
	}
}
