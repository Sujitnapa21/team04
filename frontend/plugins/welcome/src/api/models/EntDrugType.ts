/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
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
    EntDrugTypeEdges,
    EntDrugTypeEdgesFromJSON,
    EntDrugTypeEdgesFromJSONTyped,
    EntDrugTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugType
 */
export interface EntDrugType {
    /**
     * DrugTypeName holds the value of the "DrugTypeName" field.
     * @type {string}
     * @memberof EntDrugType
     */
    drugTypeName?: string;
    /**
     * 
     * @type {EntDrugTypeEdges}
     * @memberof EntDrugType
     */
    edges?: EntDrugTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDrugType
     */
    id?: number;
}

export function EntDrugTypeFromJSON(json: any): EntDrugType {
    return EntDrugTypeFromJSONTyped(json, false);
}

export function EntDrugTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugTypeName': !exists(json, 'DrugTypeName') ? undefined : json['DrugTypeName'],
        'edges': !exists(json, 'edges') ? undefined : EntDrugTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDrugTypeToJSON(value?: EntDrugType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'DrugTypeName': value.drugTypeName,
        'edges': EntDrugTypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


