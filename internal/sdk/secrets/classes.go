/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secrets

// All secret classes in this file match the secret classes used in the
// TriggerMesh web frontend.

// Secret class "aws"
const (
	secrClassAWSAccessKeyID     = "access_key_id"
	secrClassAWSSecretAccessKey = "secret_access_key"
)

// Secret class "azure_sp"
const (
	secrClassAzureSPTenantID     = "tenant_id"
	secrClassAzureSPClientID     = "client_id"
	secrClassAzureSPClientSecret = "client_secret"
)

// Secret class "basic_auth"
// https://kubernetes.io/docs/concepts/configuration/secret/#basic-authentication-secret
const (
	secrClassBasicAuthUser   = "username"
	secrClassBasicAuthPasswd = "password"
)

// Secret class "datadog"
const (
	secrClassDatadogAPIKey = "apiKey"
)

// Secret class "gcloud_service_account"
const (
	secrClassGCloudSvcAccountKey = "key.json"
)

// Secret class "github"
const (
	secrClassGitHubAccessToken   = "access_token"
	secrClassGitHubWebhookSecret = "webhook_secret"
)

// Secret class "kafka"
const (
	secrClassKafkaSASLMechanism = "sasl.mechanism"
	secrClassKafkaSASLUser      = "user"
	secrClassKafkaSASLPasswd    = "password"
	secrClassKafkaTLSCACert     = "ca.crt"
	secrClassKafkaTLSCert       = "user.crt"
	secrClassKafkaTLSKey        = "user.key"
)

// Secret class "logz"
const (
	secrClassLogzAPIToken = "token"
)

// Secret class "salesforce_oauth_jwt"
const (
	secrClassSalesforceOAuthJWTKey = "secret_key"
)

// Secret class "slack"
const (
	secrClassSlackAPIToken = "token"
)

// Secret class "slack_app"
const (
	secrClassSlackAppSignSecr = "signing_secret"
)

// Secret class "sendgrid"
const (
	secrClassSendgridAPIKey = "apiKey"
)

// Secret class "splunk_hec"
const (
	secrClassSplunkHECToken = "hec_token"
)

// Secret class "tls"
const (
	secrClassTLSCert   = "certificate"
	secrClassTLSKey    = "key"
	secrClassTLSCACert = "ca_certificate"
)

// Secret class "twilio"
const (
	secrClassTwilioSid   = "sid"
	secrClassTwilioToken = "token"
)

// Secret class "zendesk"
const (
	secrClassZendeskToken = "token"
)
