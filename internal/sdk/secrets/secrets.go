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

// secretKeySelector represents a corev1.SecretKeySelector in a format that can
// be passed to unstructured.SetNestedMap.
type secretKeySelector map[string]interface{}

// newSecretKeySelector returns a secretKeySelector with the given secret name
// and data key.
func newSecretKeySelector(name, key string) secretKeySelector {
	return secretKeySelector{
		"name": name,
		"key":  key,
	}
}

// SecretKeyRefsAWS returns secret key selectors for the "aws" secret class.
func SecretKeyRefsAWS(secretName string) (accessKeyID, secretAccessKey secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassAWSAccessKeyID),
		newSecretKeySelector(secretName, secrClassAWSSecretAccessKey)
}

// SecretKeyRefsAzureSP returns secret key selectors for the "azure_sp" secret class.
func SecretKeyRefsAzureSP(secretName string) (tenantID, clientID, clientSecret secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassAzureSPTenantID),
		newSecretKeySelector(secretName, secrClassAzureSPClientID),
		newSecretKeySelector(secretName, secrClassAzureSPClientSecret)
}

// SecretKeyRefsBasicAuth returns secret key selectors for the "basic_auth" secret class.
func SecretKeyRefsBasicAuth(secretName string) (user, passwd secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassBasicAuthUser),
		newSecretKeySelector(secretName, secrClassBasicAuthPasswd)
}

// SecretKeyRefsDatadog returns secret key selectors for the "datadog" secret class.
func SecretKeyRefsDatadog(secretName string) (key secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassDatadogAPIKey)
}

// SecretKeyRefsGCloudServiceAccount returns secret key selectors for the "gcloud_service_account" secret class.
func SecretKeyRefsGCloudServiceAccount(secretName string) (key secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassGCloudSvcAccountKey)
}

// SecretKeyRefsGitHub returns secret key selectors for the "github" secret class.
func SecretKeyRefsGitHub(secretName string) (accessToken, webhookSecret secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassGitHubAccessToken),
		newSecretKeySelector(secretName, secrClassGitHubWebhookSecret)
}

// SecretKeyRefsKafka returns secret key selectors for the "kafka" secret class.
func SecretKeyRefsKafka(secretName string) (saslMech, saslUser, saslPasswd, caCert, cert, key secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassKafkaSASLMechanism),
		newSecretKeySelector(secretName, secrClassKafkaSASLUser),
		newSecretKeySelector(secretName, secrClassKafkaSASLPasswd),
		newSecretKeySelector(secretName, secrClassKafkaTLSCACert),
		newSecretKeySelector(secretName, secrClassKafkaTLSCert),
		newSecretKeySelector(secretName, secrClassKafkaTLSKey)
}

// SecretKeyRefsLogz returns secret key selectors for the "logz" secret class.
func SecretKeyRefsLogz(secretName string) (apiToken secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassLogzAPIToken)
}

// SecretKeyRefsSalesforceOAuthJWT returns secret key selectors for the "salesforce_oauth_jwt" secret class.
func SecretKeyRefsSalesforceOAuthJWT(secretName string) (key secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassSalesforceOAuthJWTKey)
}

// SecretKeyRefsSendgrid returns secret key selectors for the "sendgrid" secret class.
func SecretKeyRefsSendgrid(secretName string) (apiKey secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassSendgridAPIKey)
}

// SecretKeyRefsSlack returns secret key selectors for the "slack" secret class.
func SecretKeyRefsSlack(secretName string) (apiToken secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassSlackAPIToken)
}

// SecretKeyRefsSlackApp returns secret key selectors for the "slack_app" secret class.
func SecretKeyRefsSlackApp(secretName string) (signSecr secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassSlackAppSignSecr)
}

// SecretKeyRefsSplunkHEC returns secret key selectors for the "splunk_hec" secret class.
func SecretKeyRefsSplunkHEC(secretName string) (hecToken secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassSplunkHECToken)
}

// SecretKeyRefsTLS returns secret key selectors for the "tls" secret class.
func SecretKeyRefsTLS(secretName string) (cert, key, caCert secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassTLSCert),
		newSecretKeySelector(secretName, secrClassTLSKey),
		newSecretKeySelector(secretName, secrClassTLSCACert)
}

// SecretKeyRefsTwilio returns secret key selectors for the "twilio" secret class.
func SecretKeyRefsTwilio(secretName string) (sid, token secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassTwilioSid),
		newSecretKeySelector(secretName, secrClassTwilioToken)
}

// SecretKeyRefsZendesk returns secret key selectors for the "zendesk" secret class.
func SecretKeyRefsZendesk(secretName string) (token secretKeySelector) {
	return newSecretKeySelector(secretName, secrClassZendeskToken)
}
