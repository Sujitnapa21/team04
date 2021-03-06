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
    EntPatientEdges,
    EntPatientEdgesFromJSON,
    EntPatientEdgesFromJSONTyped,
    EntPatientEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatient
 */
export interface EntPatient {
    /**
     * Address holds the value of the "Address" field.
     * @type {string}
     * @memberof EntPatient
     */
    address?: string;
    /**
     * Allergic holds the value of the "Allergic" field.
     * @type {string}
     * @memberof EntPatient
     */
    allergic?: string;
    /**
     * Congenital holds the value of the "Congenital" field.
     * @type {string}
     * @memberof EntPatient
     */
    congenital?: string;
    /**
     * Idcard holds the value of the "Idcard" field.
     * @type {string}
     * @memberof EntPatient
     */
    idcard?: string;
    /**
     * PatientName holds the value of the "PatientName" field.
     * @type {string}
     * @memberof EntPatient
     */
    patientName?: string;
    /**
     * 
     * @type {EntPatientEdges}
     * @memberof EntPatient
     */
    edges?: EntPatientEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatient
     */
    id?: number;
}

export function EntPatientFromJSON(json: any): EntPatient {
    return EntPatientFromJSONTyped(json, false);
}

export function EntPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'Address') ? undefined : json['Address'],
        'allergic': !exists(json, 'Allergic') ? undefined : json['Allergic'],
        'congenital': !exists(json, 'Congenital') ? undefined : json['Congenital'],
        'idcard': !exists(json, 'Idcard') ? undefined : json['Idcard'],
        'patientName': !exists(json, 'PatientName') ? undefined : json['PatientName'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPatientToJSON(value?: EntPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Address': value.address,
        'Allergic': value.allergic,
        'Congenital': value.congenital,
        'Idcard': value.idcard,
        'PatientName': value.patientName,
        'edges': EntPatientEdgesToJSON(value.edges),
        'id': value.id,
    };
}


