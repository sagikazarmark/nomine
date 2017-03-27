import { NominePage } from './app.po';

describe('nomine App', function () {
    let page: NominePage;

    beforeEach(() => {
        page = new NominePage();
    });

    it('should display the headline', () => {
        page.navigateTo();
        expect(page.getHeaderText()).toEqual('Nomine');
    });

    it('should display the subtitle', () => {
        page.navigateTo();
        expect(page.getSubtitleText()).toEqual('Your desired name');
    });
});
