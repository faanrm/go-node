import User from "../models/UserModels.js";
import jwt from 'jsonwebtoken'
import bcrypt from "bcrypt"

export const loginUser = async (req, res) => {
    const { email, password } = req.body
    try {
        const user = await User.findOne({ where: { email } })
        if (!user) {
            res.status(404).json({ message: "User not found" })
        }

        const isPassword = await bcrypt.compare(password, user.password)
        if (!isPassword) {
            return res.status(401).json({ message: 'Invalid password' })
        }

        //create token
        const token = jwt.sign({ userId: user.user_id }, process.env.TOKEN, { expiresIn: '1h' });
        // Send token in response
        res.status(200).json({ token: token, data: user });
    } catch (error) {
        console.error('Error authenticating user:', error);
        res.status(500).json({ message: 'Internal server error' });
    }
}