/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Error } from '../models/Error';
import type { User } from '../models/User';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class DefaultService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * Get's the current logged-in user
     * @returns User user response
     * @returns Error error
     * @throws ApiError
     */
    public getUser(): CancelablePromise<User | Error> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/user',
        });
    }
    /**
     * Creates a new user
     * @param requestBody
     * @returns User Creates a user
     * @returns Error error
     * @throws ApiError
     */
    public postSignup(
        requestBody: {
            name?: string;
            email?: string;
        },
    ): CancelablePromise<User | Error> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/signup',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
