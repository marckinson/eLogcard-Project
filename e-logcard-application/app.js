'use strict';
let logger=require("./custom_node_modules/logging/fileAppender.js")();
let express = require('express');
let path = require('path');
let userManagement=require("./custom_node_modules/express-middlewares/routers/user-management.js")
let bodyParser = require('body-parser');
let hfcUtil=require("./custom_node_modules/utils/hfcUtil.js");
let database=require("./custom_node_modules/database/mongo.js");
let logcardManager=require("./custom_node_modules/express-middlewares/routers/logcard-management.js");
let chainStatRouter=require("./custom_node_modules/express-middlewares/routers/chain-stat.js");

//APPLICATION

database.init()
.then(hfcUtil.initializeContext)
.then(userManagement.init)
.then(function(){
	//Initialisation de l application
	let app = express();
	app.use(bodyParser.json());
	app.use(bodyParser.urlencoded({ extended: false }));
	
	let tokenManager=require("./custom_node_modules/express-middlewares/security/token-manager.js")
	app.use(tokenManager.jwtMiddleware);
	
	app.use(userManagement.router);
	app.use("/logcard",logcardManager.router);
	app.use("/stats",chainStatRouter.router);
	
	app.listen(3000,function () {
		console.log('Application started');
	});
	
});