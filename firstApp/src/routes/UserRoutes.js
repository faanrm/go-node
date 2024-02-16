import express from "express"
import { getAllUser , createUser , deleteUser , updateUser} from "../controllers/UserControllers.js"

const router = express.Router()

router.get('/allUser',getAllUser)
router.post('/createUser',createUser)
router.delete('/deleteUser/:id',deleteUser)
router.put("/updateUser/:id",updateUser)

export default router