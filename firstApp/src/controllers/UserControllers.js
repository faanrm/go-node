import User from "../models/UserModels.js";
import jwt from "jsonwebtoken"
import bcrypt from "bcrypt"
export const getAllUser = async (req, res) => {
  try {
    const user = User.findAll()
    if (user) {
      return res.status(401).json({ msg: 'list of all users', data: user })
    }
    else {
      return res.status(404).json({ msg: 'np' })
    }

  } catch (error) {
    res.status(500).json({ message: error.message })
  }
}
export const createUser = async (req, res) => {
  const { name, email, password, confirmPassword } = req.body;
  try {
    // Check if passwords match
    if (password !== confirmPassword) {
      return res.status(400).json({ message: 'Passwords do not match' });
    }
    // Check if user with the provided email already exists
    const existingUser = await User.findOne({ where: { email } });
    if (existingUser) {
      return res.status(400).json({ message: 'User with this email already exists' });
    }
    // Hash the password
    const hashedPassword = await bcrypt.hash(password, 10);

    // Create the user
    const newUser = await User.create({ name, email, password: hashedPassword });

    // User created successfully, generate token
    const token = jwt.sign({ userId: newUser.user_id }, process.env.TOKEN, { expiresIn: '1h' });

    // Send token in response
    res.status(201).json({ token });
  } catch (error) {
    console.error('Error creating user:', error);
    res.status(500).json({ message: 'Internal server error' });
  }
}

export const deleteUser = async (req, res) => {
  const { id } = req.params
  try {
    const deletedUser = await User.destroy({
      where: { user_id: id }
    })
    if (deletedUser) {
      res.status(204).json({ message: "user deleted" }).send();
    } else {
      res.status(404).json({ message: 'User not found ' });
    }
  } catch (error) {
    res.status(500).json({ message: error.message })
  }
}
export const updateUser = async (req, res) => {
  const { id } = req.params;
  try {
    const [updated] = await User.update(req.body, {
      where: { user_id: id }
    });
    if (updated) {
      const updatedUser = await User.findByPk(id);
      res.status(200).json(updatedUser);
    } else {
      res.status(404).json({ message: 'User not found' });
    }
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};