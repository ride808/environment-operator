package translator

import (
	"testing"

	"github.com/pearsontechnology/environment-operator/pkg/bitesize"
)

func TestThirdPartyResource(t *testing.T) {
	w := BuildKubeMapper()
	w.BiteService.Type = "mysql"
	w.BiteService.Version = "5.6"

	tpr, _ := w.ThirdPartyResource()

	if tpr.Kind != "Mysql" {
		t.Errorf("tpr kind error. Expected: Mysql, got: %s", tpr.Kind)
	}

	if tpr.Spec.Version != w.BiteService.Version {
		t.Errorf("tpr version error: Expected: %s, got: %s", w.BiteService.Version, tpr.Spec.Version)
	}
}

func TestTranslatorIngressLabels(t *testing.T) {
	t.Run("ssl label", testTranslatorIngressSSl)
	t.Run("httpsBackend label", testTranslatorIngressHTTPSBackend)
	t.Run("httpsOnly label", testTranslatorIngressHTTPSOnly)
}

func testTranslatorIngressSSl(t *testing.T) {
	w := BuildKubeMapper()
	w.BiteService.Ssl = "true"

	ingress, _ := w.Ingress()

	if ingress.Labels["ssl"] != "true" {
		t.Errorf("Unexpected ingress ssl value: %+v", ingress.Labels["ssl"])
	}
}

func testTranslatorIngressHTTPSBackend(t *testing.T) {
	w := BuildKubeMapper()
	w.BiteService.HTTPSBackend = "true"

	ingress, _ := w.Ingress()

	if ingress.Labels["httpsBackend"] != "true" {
		t.Errorf("Unexpected ingress httpsBackend value: %+v", ingress.Labels["httpsBackend"])
	}
}

func testTranslatorIngressHTTPSOnly(t *testing.T) {
	w := BuildKubeMapper()
	w.BiteService.HTTPSOnly = "true"

	ingress, _ := w.Ingress()

	if ingress.Labels["httpsOnly"] != "true" {
		t.Errorf("Unexpected ingress httpsOnly value: %+v", ingress.Labels["httpsOnly"])
	}
}

func BuildKubeMapper() *KubeMapper {
	m := &KubeMapper{
		BiteService: &bitesize.Service{
			Name:  "test",
			Ports: []int{80},
		},
		Namespace: "testns",
	}
	m.Config.Project = "project"
	m.Config.DockerRegistry = "registry"
	return m
}