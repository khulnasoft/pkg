// Copyright 2019 Jason Ertel (jertel). All rights reserved.
// Copyright 2020-2023 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package modules

import (
  "github.com/khulnasoft/devsec-soc/module"
  "github.com/khulnasoft/devsec-soc/server"
  "github.com/khulnasoft/devsec-soc/server/modules/elastic"
  "github.com/khulnasoft/devsec-soc/server/modules/elasticcases"
  "github.com/khulnasoft/devsec-soc/server/modules/filedatastore"
  "github.com/khulnasoft/devsec-soc/server/modules/generichttp"
  "github.com/khulnasoft/devsec-soc/server/modules/influxdb"
  "github.com/khulnasoft/devsec-soc/server/modules/kratos"
  "github.com/khulnasoft/devsec-soc/server/modules/sostatus"
  "github.com/khulnasoft/devsec-soc/server/modules/statickeyauth"
  "github.com/khulnasoft/devsec-soc/server/modules/staticrbac"
  "github.com/khulnasoft/devsec-soc/server/modules/thehive"
)

func BuildModuleMap(srv *server.Server) map[string]module.Module {
  moduleMap := make(map[string]module.Module)
  moduleMap["filedatastore"] = filedatastore.NewFileDatastore(srv)
  moduleMap["httpcase"] = generichttp.NewHttpCase(srv)
  moduleMap["influxdb"] = influxdb.NewInfluxDB(srv)
  moduleMap["kratos"] = kratos.NewKratos(srv)
  moduleMap["elastic"] = elastic.NewElastic(srv)
  moduleMap["elasticcases"] = elasticcases.NewElasticCases(srv)
  moduleMap["sostatus"] = sostatus.NewSoStatus(srv)
  moduleMap["statickeyauth"] = statickeyauth.NewStaticKeyAuth(srv)
  moduleMap["staticrbac"] = staticrbac.NewStaticRbac(srv)
  moduleMap["thehive"] = thehive.NewTheHive(srv)
  return moduleMap
}
