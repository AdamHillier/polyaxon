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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1ModelRegistry from './V1ModelRegistry';

/**
 * The V1ListModelRegistryResponse model module.
 * @module model/V1ListModelRegistryResponse
 * @version 1.1.4
 */
class V1ListModelRegistryResponse {
    /**
     * Constructs a new <code>V1ListModelRegistryResponse</code>.
     * @alias module:model/V1ListModelRegistryResponse
     */
    constructor() { 
        
        V1ListModelRegistryResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1ListModelRegistryResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ListModelRegistryResponse} obj Optional instance to populate.
     * @return {module:model/V1ListModelRegistryResponse} The populated <code>V1ListModelRegistryResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ListModelRegistryResponse();

            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
            if (data.hasOwnProperty('results')) {
                obj['results'] = ApiClient.convertToType(data['results'], [V1ModelRegistry]);
            }
            if (data.hasOwnProperty('previous')) {
                obj['previous'] = ApiClient.convertToType(data['previous'], 'String');
            }
            if (data.hasOwnProperty('next')) {
                obj['next'] = ApiClient.convertToType(data['next'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} count
 */
V1ListModelRegistryResponse.prototype['count'] = undefined;

/**
 * @member {Array.<module:model/V1ModelRegistry>} results
 */
V1ListModelRegistryResponse.prototype['results'] = undefined;

/**
 * @member {String} previous
 */
V1ListModelRegistryResponse.prototype['previous'] = undefined;

/**
 * @member {String} next
 */
V1ListModelRegistryResponse.prototype['next'] = undefined;






export default V1ListModelRegistryResponse;
