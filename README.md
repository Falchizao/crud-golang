# api desenvolvida para a disciplina de web avançada

Api-Rest de uma aplicação CRUD em GO.
Persistência de dados realizada no PostgreSQL.

Para estruturar o projeto, deve-se iniciar o go.mod e go.sum.

- go.mod: 
    cria o modulo para exportar o projeto como package, apontando para a versão do go.

- go.sum:
    lista os packages dependencies do projeto

No arquivo .env se encontra algumas variáveis constantes para gerenciamento do projeto, adicionando
ele no .gitignore, não será enviada suas configurações privadas do projeto.

- Utilizado o fiber para handle das rotas da aplicação

- Utilizado o gorm como ORM para mapeamento das entidades

- Utilizado o godotenv para busca das configurações em arquivos locais

- Na arquivo main, as routes foram criadas e realizado um bind ao respectivo controller
ex: 	
  api := app.Group("/api")
	api.Post("/createUser", r.CreateUser)
	api.Delete("/deleteUser/:id", r.DeleteUser)
	api.Get("/get_user/:id", r.GetUserByID)
	api.Get("/users", r.GetUsers)
	api.Put("/update_user/:id", r.UpdateUser)

- No arquivo main é necessário instanciar um router gin e iniciar nossas rotas, assim ao inicializar o projeto, já é possível visualizar os endpoints

- É necessário bindar uma porta livre para o fiber, caso contrário, a aplicação não será inicializada


