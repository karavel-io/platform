# Copyright 2021 MIKAMAI s.r.l
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.15 as builder

WORKDIR /cli

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN ls
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build-simple

FROM alpine:3.13

WORKDIR /karavel

COPY --from=builder /cli/bin/karavel /usr/local/bin/karavel

ENTRYPOINT [ "/usr/local/bin/karavel" ]
CMD [ "help" ]
