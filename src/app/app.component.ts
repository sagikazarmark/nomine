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
        {
            "icon": "fa fa-twitter",
            "id": "twitter",
            "name": "Twitter",
        },
    ];

    results: Result = {
        "github": -1,
        "twitter": -1,
    };

    checkInProgress = false;

    constructor (private nomineService: NomineService) {}

    check() {
        if (this.name.length == 0) {
            return;
        }

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
