import { Component } from '@angular/core';

import { NomineService } from './nomine.service';

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
        {
            "icon": "fa fa-docker",
            "id": "docker",
            "name": "Docker",
        },
        {
            "id": "com_domain",
            "name": ".com domain",
        },
        {
            "id": "org_domain",
            "name": ".org domain",
        },
        {
            "id": "io_domain",
            "name": ".io domain",
        },
        {
            "id": "net_domain",
            "name": ".net domain",
        },
    ];

    results = {
        "github": -1,
        "twitter": -1,
        "docker": -1,
        "com_domain": -1,
        "org_domain": -1,
        "io_domain": -1,
        "net_domain": -1,
    };

    checksInProgress = 0;

    constructor (private nomineService: NomineService) {}

    check() {
        if (this.name.length == 0) {
            return;
        }

        for (var i = 0; i < this.services.length; i++) {
            this.checksInProgress++;

            var service = this.services[i].id;
            this.results[service] = -1;

            this.nomineService.check(this.name, service).subscribe(
                this.processor(service),
                //error => this.errorMessage = <any>error
            );
        }
    }

    processor(service: string) {
        return (result) => {
            this.checksInProgress--;

            this.results[service] = result;
        };
    }
}