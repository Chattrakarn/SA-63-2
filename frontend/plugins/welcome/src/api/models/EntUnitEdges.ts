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
    EntDrug,
    EntDrugFromJSON,
    EntDrugFromJSONTyped,
    EntDrugToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUnitEdges
 */
export interface EntUnitEdges {
    /**
     * Drugs holds the value of the drugs edge.
     * @type {Array<EntDrug>}
     * @memberof EntUnitEdges
     */
    drugs?: Array<EntDrug>;
}

export function EntUnitEdgesFromJSON(json: any): EntUnitEdges {
    return EntUnitEdgesFromJSONTyped(json, false);
}

export function EntUnitEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUnitEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugs': !exists(json, 'drugs') ? undefined : ((json['drugs'] as Array<any>).map(EntDrugFromJSON)),
    };
}

export function EntUnitEdgesToJSON(value?: EntUnitEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'drugs': value.drugs === undefined ? undefined : ((value.drugs as Array<any>).map(EntDrugToJSON)),
    };
}


