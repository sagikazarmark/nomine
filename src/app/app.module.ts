import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { CommonModule } from '@angular/common';

import { MaterialModule } from '@angular/material';

import 'hammerjs';

import { AppComponent } from './app.component';

import { NomineService } from './nomine.service';

@NgModule({
    declarations: [
        AppComponent
    ],
    imports: [
        BrowserModule,
        CommonModule,
        FormsModule,
        HttpModule,
        MaterialModule
    ],
    providers: [NomineService],
    bootstrap: [AppComponent]
})
export class AppModule { }
