describe('chef datafeed', () => {

  const name = 'cytest',
        url = 'http://test.com',
        username = 'admin',
        password = 'password',
        tokenType = 'TestType',
        token = 'behwveh3238238=',
        endpoint = 'https://test.com',
        bucketName = 'bucket',
        accessKey = 'access_key',
        secretKey = 'secret_key';

  const minioBucket = 'mybucket',
        minioSecret = 'minioadmin',
        minioAccess = 'minioadmin',
        minioUrl = 'http://127.0.0.1:9000';


  before(() => {
    cy.adminLogin('/settings').then(() => {
        const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
        cy.get('app-welcome-modal').invoke('hide');
        cy.restoreStorage();
        cy.get('body').type('feat');
        cy.get('.title').contains('Chef Automate Data Feed').parent().parent()
          .find('.onoffswitch').click();
        cy.get('chef-button').contains('Close').click();
        cy.reload();
        cy.contains('know').click();
        cy.contains('Data Feeds').click();
        cy.get('app-notification.error chef-icon').click({ multiple: true });
    });
  });

  beforeEach(() => {
    cy.restoreStorage();
  });

  afterEach(() => {
      cy.saveStorage();
  });

  describe ('chef data feed page', () => {
    const reusableDate = Date.now();

    it('check if clicking on new integration button opens up the slider', () => {
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=interation-menu]').should('be.visible');
      cy.get('[data-cy=close-feed-button]').click();
    });

    it('check if clicking on a interation opens the form', () => {
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=ServiceNow]').click();
      cy.get('[data-cy=data-feed-create-form]').should('be.visible');
      cy.get('[data-cy=close-feed-button]').click();
    });

    it('create data feed service now', () => {
      const date = Date.now();
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=ServiceNow]').click();
      cy.get('[data-cy=add-name]').type(name + date);
      cy.get('[data-cy=add-url]').type(url);
      cy.get('[data-cy=select-auth-type]').click();
      cy.get('[data-cy=select-username-password]').click();
      cy.get('[data-cy=add-username]').type(username);
      cy.get('[data-cy=add-password]').type(password);
      cy.get('[data-cy=add-button]').click();
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click();
      cy.contains('Data Feeds').click();
      cy.get('chef-table chef-tbody chef-td').contains('cytest' + date).should('exist');
    });

    it('create data feed splunk', () => {
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=Splunk]').click();
      cy.get('[data-cy=add-name]').type(name + reusableDate);
      cy.get('[data-cy=add-url]').type(url);
      cy.get('[data-cy=add-token]').type(token);
      cy.get('[data-cy=add-button]').click();
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click();
      cy.contains('Data Feeds').click();
      cy.get('chef-table chef-tbody chef-td').contains('cytest' + reusableDate).should('exist');
    });

    it('create data feed error', () => {
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=Splunk]').click();
      cy.get('[data-cy=add-name]').type(name + reusableDate);
      cy.get('[data-cy=add-url]').type(url);
      cy.get('[data-cy=add-token]').type(token);
      cy.get('[data-cy=add-button]').click();
      cy.get('app-data-feed-create').scrollTo('top');
      cy.get('.data-feed-slider app-notification.error').should('be.visible');
      cy.get('.data-feed-slider app-notification.error chef-icon').click();
      cy.get('[data-cy=close-feed-button]').click();
    });

    it('create data feed with changed token type', () => {
      const date = Date.now();
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=Splunk]').click();
      cy.get('[data-cy=add-name]').type(name + date);
      cy.get('[data-cy=add-url]').type(url);
      cy.get('[data-cy=toggle-type]').click();
      cy.get('[data-cy=add-token-type]').clear();
      cy.get('[data-cy=add-token-type]').type(tokenType);
      cy.get('[data-cy=add-token]').type(token);
      cy.get('[data-cy=add-button]').click();
      cy.get('app-notification.info').should('be.visible');
      cy.get('app-notification.info chef-icon').click();
      cy.contains('Data Feeds').click();
      cy.get('chef-table chef-tbody chef-td').contains('cytest' + date).should('exist');
    });

    it('test error in data feed', () => {
      const date = Date.now();
      cy.get('[data-cy=create-data-feed]').click();
      cy.get('[data-cy=Splunk]').click();
      cy.get('[data-cy=add-name]').type(name + date);
      cy.get('[data-cy=add-url]').type(url);
      cy.get('[data-cy=add-token]').type(token);
      cy.get('[data-cy=test-button]').click();
      cy.get('app-data-feed-create').scrollTo('top');
      cy.get('.data-feed-slider app-notification.error').should('be.visible');
      cy.get('.data-feed-slider app-notification.error chef-icon').click();
      cy.get('[data-cy=close-feed-button]').click();
    });
  });
});
