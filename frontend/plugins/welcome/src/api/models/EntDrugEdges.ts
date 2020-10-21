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
    EntForm,
    EntFormFromJSON,
    EntFormFromJSONTyped,
    EntFormToJSON,
    EntUnit,
    EntUnitFromJSON,
    EntUnitFromJSONTyped,
    EntUnitToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
    EntVolume,
    EntVolumeFromJSON,
    EntVolumeFromJSONTyped,
    EntVolumeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugEdges
 */
export interface EntDrugEdges {
    /**
     * 
     * @type {EntForm}
     * @memberof EntDrugEdges
     */
    form?: EntForm;
    /**
     * 
     * @type {EntUnit}
     * @memberof EntDrugEdges
     */
    unit?: EntUnit;
    /**
     * 
     * @type {EntUser}
     * @memberof EntDrugEdges
     */
    user?: EntUser;
    /**
     * 
     * @type {EntVolume}
     * @memberof EntDrugEdges
     */
    volume?: EntVolume;
}

export function EntDrugEdgesFromJSON(json: any): EntDrugEdges {
    return EntDrugEdgesFromJSONTyped(json, false);
}

export function EntDrugEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'form': !exists(json, 'Form') ? undefined : EntFormFromJSON(json['Form']),
        'unit': !exists(json, 'Unit') ? undefined : EntUnitFromJSON(json['Unit']),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
        'volume': !exists(json, 'Volume') ? undefined : EntVolumeFromJSON(json['Volume']),
    };
}

export function EntDrugEdgesToJSON(value?: EntDrugEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'form': EntFormToJSON(value.form),
        'unit': EntUnitToJSON(value.unit),
        'user': EntUserToJSON(value.user),
        'volume': EntVolumeToJSON(value.volume),
    };
}


