/**
 * Copyright 2016, 2017 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core_test

import (
	"fmt"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/ubiquity-docker-plugin/core"
	"github.com/IBM/ubiquity/fakes"

	"github.com/IBM/ubiquity/resources"
	"time"
	"strconv"
	"reflect"
)

var _ = Describe("Controller", func() {
	Context("on activate", func() {
		var (
			fakeClient *fakes.FakeStorageClient
			controller *core.Controller
			random     *rand.Rand
		)
		BeforeEach(func() {
			fakeClient = new(fakes.FakeStorageClient)
			backends := []string{Backend}
			controller = core.NewControllerWithClient(testLogger, fakeClient, backends)
			random = rand.New(rand.NewSource(time.Now().UnixNano()))
		})
		It("does not error when remote client activate is successful", func() {
			fakeClient.ActivateReturns(nil)
			activateResponse := controller.Activate()
			Expect(activateResponse.Implements).ToNot(Equal(nil))
			Expect(len(activateResponse.Implements)).To(Equal(1))
			Expect(activateResponse.Implements[0]).To(Equal("VolumeDriver"))
		})

		It("errors when remote client activate fails", func() {
			fakeClient.ActivateReturns(fmt.Errorf("failed to activate"))
			activateResponse := controller.Activate()
			Expect(activateResponse.Implements).ToNot(Equal(nil))
			Expect(len(activateResponse.Implements)).To(Equal(0))
		})

		Context("on successful activate", func() {
			BeforeEach(func() {
				fakeClient.ActivateReturns(nil)
				activateResponse := controller.Activate()
				Expect(activateResponse.Implements).ToNot(Equal(nil))
				Expect(len(activateResponse.Implements)).To(Equal(1))
				Expect(activateResponse.Implements[0]).To(Equal("VolumeDriver"))
			})
			Context(".Create", func() {
				It("does not error on create with valid opts", func() {
					fakeClient.CreateVolumeReturns(nil)
					createRequest := resources.CreateVolumeRequest{Name: "dockerVolume1", Backend:Backend, Opts: map[string]interface{}{"Filesystem": "gpfs1"}}
					createResponse := controller.Create(createRequest)
					Expect(createResponse.Err).To(Equal(""))
					Expect(fakeClient.CreateVolumeCallCount()).To(Equal(1))
					name := fakeClient.CreateVolumeArgsForCall(0).Name
					Expect(name).To(Equal("dockerVolume1"))
				})
				It("does error on create when plugin fails to create dockerVolume", func() {
					fakeClient.CreateVolumeReturns(fmt.Errorf("Spectrum plugin internal error"))
					createRequest := resources.CreateVolumeRequest{Name: "dockerVolume1",Backend:Backend, Opts: map[string]interface{}{"Filesystem": "gpfs1"}}
					createResponse := controller.Create(createRequest)
					Expect(createResponse.Err).To(Equal("Spectrum plugin internal error"))
				})
			})
			Context(".Remove", func() {
				It("does not error when existing dockerVolume name is given", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					removeRequest := resources.RemoveVolumeRequest{Name: "dockerVolume1"}
					removeResponse := controller.Remove(removeRequest)
					Expect(removeResponse.Err).To(Equal(""))
				})
				It("error when remove dockerVolume returns an error", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					fakeClient.RemoveVolumeReturns(fmt.Errorf("error removing volume"))
					removeRequest := resources.RemoveVolumeRequest{Name: "dockerVolume1"}
					removeResponse := controller.Remove(removeRequest)
					Expect(removeResponse.Err).To(Equal("error removing volume"))
					Expect(fakeClient.RemoveVolumeCallCount()).To(Equal(1))
				})
			})
			Context(".List", func() {
				It("does not error when volumes exist", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"}
					var dockerVolumes []resources.Volume
					dockerVolumes = append(dockerVolumes, dockerVolume)
					fakeClient.ListVolumesReturns(dockerVolumes, nil)
					listResponse := controller.List()
					Expect(listResponse.Err).To(Equal(""))
					Expect(listResponse.Volumes).ToNot(Equal(nil))
					Expect(len(listResponse.Volumes)).To(Equal(1))
				})
				It("does not error when no volumes exist", func() {
					var dockerVolumes []resources.Volume
					fakeClient.ListVolumesReturns(dockerVolumes, nil)
					listResponse := controller.List()
					Expect(listResponse.Err).To(Equal(""))
					Expect(listResponse.Volumes).ToNot(Equal(nil))
					Expect(len(listResponse.Volumes)).To(Equal(0))
				})
				It("errors when client fails to list dockerVolumes", func() {
					fakeClient.ListVolumesReturns(nil, fmt.Errorf("failed to list volumes"))
					listResponse := controller.List()
					Expect(listResponse.Err).To(Equal("failed to list volumes"))
				})
			})
			Context(".Get", func() {
				It("does not error when volume exist", func() {
					config := make(map[string]interface{})
					config["mountpoint"] = "some-mountpoint"
					fakeClient.GetVolumeConfigReturns(config, nil)
					getRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					getResponse := controller.Get(getRequest)
					Expect(getResponse.Err).To(Equal(""))
					Expect(getResponse.Volume).ToNot(Equal(nil))
					Expect(getResponse.Volume["Name"]).To(Equal("dockerVolume1"))
				})
				It("errors when list dockerVolume returns an error", func() {
					fakeClient.GetVolumeConfigReturns(nil, fmt.Errorf("failed listing volume"))
					getRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					getResponse := controller.Get(getRequest)
					Expect(getResponse.Err).To(Equal("failed listing volume"))
				})
				It("get Status from backend", func() {
					keyStr, valStr := "key", "val"
					num := 1 + random.Intn(10)
					config := make(map[string]interface{})
					for index := 0; index < num; index++ {
						indexStr := strconv.Itoa(index)
						config[keyStr+indexStr] = valStr + indexStr
					}
					fakeClient.GetVolumeConfigReturns(config, nil)
					getRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					getResponse := controller.Get(getRequest)
					Expect(getResponse.Err).To(Equal(""))
					Expect(getResponse.Volume["Name"]).To(Equal("dockerVolume1"))
					eq := reflect.DeepEqual(getResponse.Volume["Status"], config)
					Expect(eq).To(Equal(true))
				})
			})
			Context(".Path", func() {
				It("does not error when volume exists and is mounted", func() {
					config := make(map[string]interface{})
					config["mountpoint"] = "some-mountpoint"
					fakeClient.GetVolumeConfigReturns(config, nil)
					pathRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					pathResponse := controller.Path(pathRequest)
					Expect(pathResponse.Err).To(Equal(""))
					Expect(pathResponse.Mountpoint).To(Equal("some-mountpoint"))
				})
				It("errors when volume exists but is not mounted", func() {
					config := make(map[string]interface{})
					fakeClient.GetVolumeConfigReturns(config, nil)
					pathRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					pathResponse := controller.Path(pathRequest)
					Expect(pathResponse.Err).To(Equal("volume not mounted"))
				})
				It("errors when list dockerVolume returns an error", func() {

					fakeClient.GetVolumeConfigReturns(nil, fmt.Errorf("failed listing volume"))
					pathRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					pathResponse := controller.Path(pathRequest)
					Expect(pathResponse.Err).To(Equal("failed listing volume"))
				})
				It("errors when volume does not exist", func() {
					fakeClient.GetVolumeConfigReturns(nil, fmt.Errorf("volume does not exist"))
					pathRequest := resources.GetVolumeConfigRequest{Name: "dockerVolume1"}
					pathResponse := controller.Path(pathRequest)
					Expect(pathResponse.Err).To(Equal("volume does not exist"))
				})
			})
			Context(".Mount", func() {
				It("does not error when volume exists and is not currently mounted", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					fakeClient.AttachReturns("some-mountpath", nil)
					mountRequest := resources.AttachRequest{Name: "dockerVolume1"}
					mountResponse := controller.Mount(mountRequest)
					Expect(mountResponse.Err).To(Equal(""))
					Expect(mountResponse.Mountpoint).To(Equal("some-mountpath"))
					Expect(fakeClient.AttachCallCount()).To(Equal(1))
				})

				It("errors when volume exists and linking dockerVolume errors", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					fakeClient.AttachReturns("", fmt.Errorf("failed to link volume"))
					mountRequest := resources.AttachRequest{Name: "dockerVolume1"}
					mountResponse := controller.Mount(mountRequest)
					Expect(mountResponse.Err).To(Equal("failed to link volume"))
				})
			})
			Context(".Unmount", func() {
				It("does not error when volume exists and is currently mounted", func() {

					dockerVolume := resources.Volume{Name: "dockerVolume1"} //, Mountpoint: "some-mountpoint"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					unmountRequest := resources.DetachRequest{Name: "dockerVolume1"}
					unmountResponse := controller.Unmount(unmountRequest)
					Expect(unmountResponse.Err).To(Equal(""))
				})

				It("errors when volume exists and un-linking dockerVolume errors", func() {
					dockerVolume := resources.Volume{Name: "dockerVolume1"} //, Mountpoint: "some-mountpoint"}
					fakeClient.GetVolumeReturns(dockerVolume, nil)
					fakeClient.DetachReturns(fmt.Errorf("failed to unlink volume"))
					unmountRequest := resources.DetachRequest{Name: "dockerVolume1"}
					unmountResponse := controller.Unmount(unmountRequest)
					Expect(unmountResponse.Err).To(Equal("failed to unlink volume"))
				})
			})
		})
	})
})
