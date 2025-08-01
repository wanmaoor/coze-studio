/*
 * Copyright 2025 coze-dev Authors
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

package template

import (
	"context"

	"gorm.io/gorm"

	"github.com/coze-dev/coze-studio/backend/domain/template/repository"
	"github.com/coze-dev/coze-studio/backend/infra/contract/idgen"
	"github.com/coze-dev/coze-studio/backend/infra/contract/storage"
)

type ServiceComponents struct {
	DB      *gorm.DB
	IDGen   idgen.IDGenerator
	Storage storage.Storage
}

func InitService(ctx context.Context, components *ServiceComponents) *ApplicationService {

	tRepo := repository.NewTemplateDAO(components.DB, components.IDGen)

	ApplicationSVC.templateRepo = tRepo
	ApplicationSVC.storage = components.Storage

	return ApplicationSVC
}
