// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libmonteur

// Supported variables keys in key:value variables placholders.
//
// It is used in every toml config file inside setup/program/ config directory
// for placeholding variable elements in the fields' value.
const (
	VAR_ARCH                      = "Arch"
	VAR_ARCHIVE                   = "Archive"
	VAR_APP                       = "App"
	VAR_BASE                      = "BaseDir"
	VAR_BUILD                     = "BuildDir"
	VAR_BIN                       = "BinDir"
	VAR_CFG                       = "ConfigDir"
	VAR_CHANGELOG_ENTRIES         = "ChangelogEntries"
	VAR_COMPUTE                   = "ComputeSystem"
	VAR_DATA                      = "DataDir"
	VAR_DOC                       = "DocsDir"
	VAR_FORMAT                    = "Format"
	VAR_HOME                      = "HomeDir"
	VAR_LOG                       = "LogDir"
	VAR_METHOD                    = "Method"
	VAR_OS                        = "OS"
	VAR_PACKAGE                   = "PackageDir"
	VAR_ROOT                      = "RootDir"
	VAR_SOURCE                    = "Source"
	VAR_SECRETS                   = "Secrets"
	VAR_TARGET                    = "Target"
	VAR_TMP                       = "WorkingDir"
	VAR_URL                       = "URL"
	VAR_PACKAGE_NAME              = "PkgName"
	VAR_PACKAGE_VERSION           = "PkgVersion"
	VAR_PACKAGE_VERSION_DIGIT_LED = "PkgVersionDigitLed"
	VAR_PACKAGE_OS                = "PkgOS"
	VAR_PACKAGE_ARCH              = "PkgArch"
	VAR_SOURCE_ARCH               = "SourceArch"
	VAR_SOURCE_COMPUTE            = "SourceCompute"
	VAR_SOURCE_OS                 = "SourceOS"
	VAR_RELEASE                   = "ReleaseDir"
	VAR_TIMESTAMP                 = "Timestamp"
)
