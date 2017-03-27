import { browser, element, by } from 'protractor';

export class NominePage {
    navigateTo() {
        return browser.get('/');
    }

    getHeaderText() {
        return element(by.css('app header h1')).getText();
    }

    getSubtitleText() {
        return element(by.css('app header p')).getText();
    }
}
