describe('Тест регистрации и авторизации', () => {
  const user = {
    email: 'testuser@example.com',
    password: 'TestPassword123',
    firstName: 'Тест',
    lastName: 'Пользователь'
  };

  before(() => {
    cy.visit('https://localhost:8080');
  });

  it('Регистрируем тестового пользователя', () => {
    cy.get('button').contains('Создать').click();

    cy.get('input#email').clear().type(user.email);
    cy.get('input#password').clear().type(user.password);
    cy.get('input#surname').clear().type(user.lastName);
    cy.get('input#name').clear().type(user.firstName);

    cy.get('button').contains('Создать').click();
  });

  it('Пробуем авторизоваться под тестовым пользователем', () => {
    cy.visit('https://localhost:8080');

    cy.get('input#email').clear().type(user.email);
    cy.get('input#password').clear().type(user.password);

    cy.get('button').contains('Войти').click();

    // Проверка успешного входа на домашнюю страницу
    cy.url().should('include', '/home');

    // Проверка наличия токена в localStorage
    cy.window().then((win) => {
      const token = win.localStorage.getItem('token');
      expect(token).to.exist;
    });

    // Проверка наличия куки auth
    cy.getCookie('auth').should('exist');
  });

  it('Нажимаем кнопку выхода', () => {
    // Предполагается, что на странице home есть кнопка "Выйти"
    cy.get('button').contains('Выйти').click();

    // Проверка, что нас вернуло на страницу авторизации
    cy.url().should('include', '#/');
    cy.get('h1').should('contain', 'Авторизация');
  });

  it('Пробуем повторно создать пользователя на ту же почту', () => {
    cy.get('button').contains('Создать').click();

    cy.get('input#email').clear().type(user.email);
    cy.get('input#password').clear().type(user.password);
    cy.get('input#surname').clear().type(user.lastName);
    cy.get('input#name').clear().type(user.firstName);

    cy.get('button').contains('Создать').click();

    // Проверка, что отображается сообщение об ошибке регистрации
    cy.get('h3').should('contain', 'Ошибка регистрации');
  });
});
