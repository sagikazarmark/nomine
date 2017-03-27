import { Component } from '@angular/core';

import { NomineService, Result } from './nomine.service';

@Component({
    selector: 'app',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    name: string;

    services = [
        {
            "icon": "fa fa-github",
            "id": "github",
            "name": "Github",
        },
    ];

    results: Result = {
        "github": -1,
    };

    checkInProgress = false;

    constructor (private nomineService: NomineService) {}

    check() {
        this.checkInProgress = true;

        var s: string[] = [];

        for (var i = 0; i < this.services.length; i++) {
            s.push(this.services[i].id);
        }

        this.nomineService.check(this.name, s).subscribe(
            results => {
                this.checkInProgress = false;
                this.results = results;
            },
            //error => this.errorMessage = <any>error
        );
    }
}
