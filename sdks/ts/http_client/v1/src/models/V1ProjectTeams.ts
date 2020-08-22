// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.8-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1ProjectTeams
 */
export interface V1ProjectTeams {
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ProjectTeams
     */
    teams?: Array<string>;
}

export function V1ProjectTeamsFromJSON(json: any): V1ProjectTeams {
    return V1ProjectTeamsFromJSONTyped(json, false);
}

export function V1ProjectTeamsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ProjectTeams {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'teams': !exists(json, 'teams') ? undefined : json['teams'],
    };
}

export function V1ProjectTeamsToJSON(value?: V1ProjectTeams | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'teams': value.teams,
    };
}


