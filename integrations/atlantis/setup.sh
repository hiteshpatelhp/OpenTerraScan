#   Copyright (C) 2022 Tenable, Inc.
#
#	  Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#		http://www.apache.org/licenses/LICENSE-2.0
#
#	  Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

#!/bin/bash
set -ex

if [[ -z "${OPENTERRASCAN_VERSION}" ]]; then
  OPENTERRASCAN_VERSION=${DEFAULT_OPENTERRASCAN_VERSION}
fi

VERSION=${OPENTERRASCAN_VERSION}

curl -LOs https://github.com/tenable/openterrascan/releases/download/v${VERSION}/openterrascan_${VERSION}_Linux_x86_64.tar.gz
mkdir /usr/local/bin/openterrascan_${VERSION}
tar -C  /usr/local/bin/openterrascan_${VERSION} -xzf openterrascan_${VERSION}_Linux_x86_64.tar.gz

mv /usr/local/bin/openterrascan_${VERSION}/openterrascan /usr/local/bin/openterrascan

rm openterrascan_${VERSION}_Linux_x86_64.tar.gz
rm -rf /usr/local/bin/openterrascan_${VERSION}/
