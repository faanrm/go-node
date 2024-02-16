import express from "express";
import cors from "cors";
import db from "./config/sql.js";
import UserRouter from "./routes/UserRoutes.js"
import dotenv from "dotenv"
import AuthRouter from "./routes/AuthRoutes.js"
dotenv.config()
const app = express();
const port = process.env.APP_PORT || 5200;

app.use(cors());
app.use(express.json())

//routes & middleware
app.get("/", (req, res) => {
  res.send("Hello World");
});
app.use(UserRouter)
app.use(AuthRouter)

//connection to the db
// Start the server only after the database connection has been established
db.sync().then(() => {
  app.listen(port, () => {
    console.log(`Server Running on Port: http://localhost:${port}`);
  });
}).catch(err => {
  console.error('Unable to connect to the database:', err);
  process.exit(-1);
});

