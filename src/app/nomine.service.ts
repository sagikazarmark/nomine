import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

import { environment } from '../environments/environment'

@Injectable()
export class NomineService {
    private url = '/check';

    constructor(private http: Http) {}

    check(name: string, service: string): Observable<number> {
        return this.http.get(environment.apiEndpoint + this.url + "/" + service + "/" + name)
            .map((res: Response) => {
                let body = res.json();
                return body.result || { };
            })
        ;
    }
}
