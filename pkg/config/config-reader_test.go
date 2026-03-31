/*
    Copyright (C) 2022 Tenable, Inc.

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

package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

var (
	testDataDir = "testdata"

	testRules = Rules{
		ScanRules: []string{"rule.1", "rule.2", "rule.3", "rule.4", "rule.5"},
		SkipRules: []string{"rule.1"},
	}

	testCategoryList = Category{List: []string{"category.1", "category.2"}}

	testNotifier = Notifier{
		NotifierType: "webhook",
		NotifierConfig: map[string]interface{}{
			"url": "testurl1",
		},
	}

	testK8sAdmControl = K8sAdmissionControl{
		Dashboard:      true,
		DeniedSeverity: highSeverity.Level,
		Categories:     testCategoryList.List,
		SaveRequests:   true,
	}

	highSeverity = Severity{Level: "high"}
)

func TestNewOpenTerraScanConfigReader(t *testing.T) {

	testPolicy := Policy{
		RepoPath: "rego-subdir",
		BasePath: "custom-path",
		RepoURL:  "https://repository/url",
		Branch:   "branch-name",
	}

	type args struct {
		fileName string
	}

	tests := []struct {
		name          string
		args          args
		want          *OpenTerraScanConfigReader
		wantErr       bool
		assertGetters bool
		Policy
		notifications map[string]Notifier
		Rules
	}{
		{
			name: "empty config file",
			args: args{
				fileName: "",
			},
			want: &OpenTerraScanConfigReader{},
		},
		{
			name: "nonexistent config file",
			args: args{
				fileName: "test",
			},
			wantErr: true,
			want:    &OpenTerraScanConfigReader{},
		},
		{
			name: "invalid config file format",
			args: args{
				fileName: "test.invalid",
			},
			wantErr: true,
			want:    &OpenTerraScanConfigReader{},
		},
		{
			name: "invalid toml config file",
			args: args{
				fileName: filepath.Join(testDataDir, "invalid.toml"),
			},
			wantErr: true,
			want:    &OpenTerraScanConfigReader{},
		},
		{
			name: "invalid yaml config file",
			args: args{
				fileName: filepath.Join(testDataDir, "invalid.toml"),
			},
			wantErr: true,
			want:    &OpenTerraScanConfigReader{},
		},
		{
			name: "valid toml config file with partial fields",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config.toml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
				},
			},
		},
		{
			name: "valid toml config file with all fields",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-all-fields.toml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:               testRules,
					Severity:            highSeverity,
					Category:            testCategoryList,
					K8sAdmissionControl: testK8sAdmControl,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid toml config file with all fields and severity defined",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-severity.toml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Severity: highSeverity,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid toml config file with all fields and categories defined",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-category.toml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Category: testCategoryList,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-all-fields.yaml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:               testRules,
					Severity:            highSeverity,
					Category:            testCategoryList,
					K8sAdmissionControl: testK8sAdmControl,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields and severity defined",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-severity.yml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Severity: highSeverity,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields and categories defined",
			args: args{
				fileName: filepath.Join(testDataDir, "openterrascan-config-category.yaml"),
			},
			want: &OpenTerraScanConfigReader{
				config: OpenTerraScanConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Category: testCategoryList,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOpenTerraScanConfigReader(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOpenTerraScanConfigReader() got error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpenTerraScanConfigReader() = got %v, want %v", got, tt.want)
			}
			if tt.assertGetters {
				if !reflect.DeepEqual(got.getPolicyConfig(), tt.Policy) || !reflect.DeepEqual(got.getNotifications(), tt.notifications) || !reflect.DeepEqual(got.getRules(), tt.Rules) {
					t.Errorf("NewOpenTerraScanConfigReader() = got config: %v, notifications: %v, rules: %v want config: %v, notifications: %v, rules: %v", got.getPolicyConfig(), got.getNotifications(), got.getRules(), tt.Policy, tt.notifications, tt.Rules)
				}
			}
		})
	}
}
