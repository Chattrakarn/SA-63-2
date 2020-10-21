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
    EntUserEdges,
    EntUserEdgesFromJSON,
    EntUserEdgesFromJSONTyped,
    EntUserEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUser
 */
export interface EntUser {
    /**
     * Email holds the value of the "Email" field.
     * @type {string}
     * @memberof EntUser
     */
    email?: string;
    /**
     * Password holds the value of the "Password" field.
     * @type {string}
     * @memberof EntUser
     */
    password?: string;
    /**
     * Username holds the value of the "Username" field.
     * @type {string}
     * @memberof EntUser
     */
    username?: string;
    /**
     * 
     * @type {EntUserEdges}
     * @memberof EntUser
     */
    edges?: EntUserEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntUser
     */
    id?: number;
}

export function EntUserFromJSON(json: any): EntUser {
    return EntUserFromJSONTyped(json, false);
}

export function EntUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'Email') ? undefined : json['Email'],
        'password': !exists(json, 'Password') ? undefined : json['Password'],
        'username': !exists(json, 'Username') ? undefined : json['Username'],
        'edges': !exists(json, 'edges') ? undefined : EntUserEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntUserToJSON(value?: EntUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Email': value.email,
        'Password': value.password,
        'Username': value.username,
        'edges': EntUserEdgesToJSON(value.edges),
        'id': value.id,
    };
}


