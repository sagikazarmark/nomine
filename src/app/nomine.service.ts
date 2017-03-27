import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

import { environment } from '../environments/environment'

@Injectable()
export class NomineService {
    private url = '/check';

    constructor(private http: Http) {}

    check(name: string, services: string[]): Observable<Result> {
        console.log(environment.apiEndpoint + this.url)
        return this.http.post(environment.apiEndpoint + this.url, { name: name, services: services })
            .map((res: Response) => {
                let body = res.json();
                return body.results || { };
            })
        ;
    }
}

export interface Result {
    [key: string]: number;
}
