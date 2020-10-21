/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntFormEdges,
    EntFormEdgesFromJSON,
    EntFormEdgesFromJSONTyped,
    EntFormEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntForm
 */
export interface EntForm {
    /**
     * FormType holds the value of the "FormType" field.
     * @type {string}
     * @memberof EntForm
     */
    formType?: string;
    /**
     * 
     * @type {EntFormEdges}
     * @memberof EntForm
     */
    edges?: EntFormEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntForm
     */
    id?: number;
}

export function EntFormFromJSON(json: any): EntForm {
    return EntFormFromJSONTyped(json, false);
}

export function EntFormFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntForm {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'formType': !exists(json, 'FormType') ? undefined : json['FormType'],
        'edges': !exists(json, 'edges') ? undefined : EntFormEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFormToJSON(value?: EntForm | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'FormType': value.formType,
        'edges': EntFormEdgesToJSON(value.edges),
        'id': value.id,
    };
}


