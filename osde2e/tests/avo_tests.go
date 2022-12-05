package tests

import (
	"context"
	"log"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
)

var crdName = ""

var _ = ginkgo.Describe("AVO Tests", func() {
	defer ginkgo.GinkgoRecover()
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err)
	}

	ginkgo.It("  CRD - "+crdName+" - exists", func() {
		apiextensions, err := clientset.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred())

		// Make sure the CRD exists
		result, err := apiextensions.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), "", v1.GetOptions{})

		if err != nil {
			log.Printf("CRD not found: %v", err.Error())
		} else {
			log.Printf("CRD found: %v", result)
		}

		Expect(err).NotTo(HaveOccurred())
	}, float64(30))

})
