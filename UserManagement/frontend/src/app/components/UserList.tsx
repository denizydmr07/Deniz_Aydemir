import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styles from './list.module.css';

// Define the user interface
interface User {
  id: string;
  username: string;
  password: string;
}


// Define the property type
interface UserListItemProps {
  user: User;
  selected: boolean;
  editMode: boolean;
  deleteMode: boolean;
  newMode: boolean;
  setEditMode: (editMode: boolean) => void;
  setDeleteMode: (deleteMode: boolean) => void;
  setNewMode: (newMode: boolean) => void;
  onClick: () => void;
}

// Define the UserListItem component
const UserListItem: React.FC<UserListItemProps> = (props) => {

  // Destructure the properties
  const { user, selected, editMode, deleteMode, newMode, setEditMode, setDeleteMode, setNewMode, onClick } = props;

  // Define the state variables for the component
  const [id, setId] = useState<string>(user.id);
  const [username, setUsername] = useState<string>(user.username);
  const [password, setPassword] = useState<string>(user.password);
  
  /**
   * Delete the user
   */
  const deleteUser = async () => {

    console.log("Deleting user");

    // Send a DELETE request
    await axios.delete(`http://localhost:8080/api/deleteUser?id=${id}`, 
    ).then(response => {

      // Handle success
      console.log(response);
      alert(response)
    }).catch(error => {

      // Handle error
      console.error('Error deleting user:', error);
    });

    // Reload the page
    window.location.reload();
  }

  /**
   * Update the user
   */
  const updateUser = async () => {
    console.log("Updating user");

    // Send a PUT request
    await axios.put(`http://localhost:8080/api/updateUser?id=${id}&username=${username}&password=${password}`, 
    ).then(response => {

      // Handle success
      console.log(response.data.message);
    }).catch(error => {

      // Handle error
      console.error('Error updating user:', error);
    });

    // Reload the page
    window.location.reload();
  }

  /**
   * Create a new user
   */
  const createNewUser = async () => {
    console.log("Creating new user");

    // Send a POST request
    await axios.post(`http://localhost:8080/api/saveUser?username=${username}&password=${password}`, 
    ).then(response => {

      // Handle success
      console.log(response.data.message);
    }).catch(error => {

      // Handle error
      console.error('Error creating new user:', error);
    });


    // Reload the page
    window.location.reload();
  }

  // Return the JSX element
  return (
    <div className={`${styles.list_item} ${selected ? styles.selected : ''}`} onClick={onClick}>
      <div>
        {}
        {editMode && (
          <div>
            <div>
              <span>Username: </span>
              <input 
                type="text" 
                value={username} 
                onChange={
                  (e) => {
                    // Handle change
                    setUsername(e.target.value);
                  }
                }
              />
            </div>
            <div>
              <span>Password: </span>
              <input 
                type="text" 
                value={password} 
                onChange={
                  (e) => {
                    // Handle change
                    setPassword(e.target.value);
                  }
                }
                />
            </div>
            <div>
              <button onClick={
                () => {
                  // Handle save
                  updateUser();
                  setEditMode(false);
                }
              }>Save</button>
              <button onClick={
                () => {
                  // Handle cancel
                  setEditMode(false);
                }
              }>Cancel</button>
            </div>
          </div>
        )}
        {
          deleteMode && (
            <div>
              <div>
                <div>Username: {user.username}</div>
                <div>Password: {user.password}</div>
              </div>
              <span>Are you sure you want to delete this user?</span>
              <button onClick={
                () => {
                  // Handle delete
                  deleteUser();
                  setDeleteMode(false);
                }
              }>Delete</button>
              <button onClick={
                () => {
                  // Handle cancel
                  setDeleteMode(false);
                }
              }>Cancel</button>
            </div>
          )
        }
        {
          newMode && (
            <div>
              <div>
                <span>Username: </span>
                <input 
                  type="text" 
                  value={username} 
                  onChange={
                    (e) => {
                      // Handle change
                      setUsername(e.target.value);
                    }
                  }
                />
              </div>
              <div>
                <span>Password: </span>
                <input 
                  type="text" 
                  value={password} 
                  onChange={
                    (e) => {
                      // Handle change
                      setPassword(e.target.value);
                    }
                  }
                  />
              </div>
              <div>
                <button onClick={
                  () => {
                    // Handle save
                    createNewUser();
                    setNewMode(false);
                  }
                }>Create</button>
                <button onClick={
                  () => {
                    // Handle cancel
                    setNewMode(false);
                  }
                }>Cancel</button>
              </div>
            </div>
          )
        }
        {!editMode && !deleteMode && !newMode && (
          <div>
            <div>Username: {user.username}</div>
            <div>Password: {user.password}</div>
          </div>
        )}
      </div>
    </div>
  );
};


// Define the UserList component
const UserList: React.FC = () => {

  // Define the state variables for the component
  const [users, setUsers] = useState<User[]>([]);
  const [editMode, setEditMode] = useState<boolean>(false);
  const [deleteMode, setDeleteMode] = useState<boolean>(false);
  const [newMode, setNewMode] = useState<boolean>(false);
  const [selectedIndex, setSelectedIndex] = useState<number>(-1);


  /**
   * Handle the edit button click
   */
  const handleEdit = () => {

    // Set the edit mode
    setEditMode(true);
    setDeleteMode(false);
    setNewMode(false);
  };


  /**
   * Handle the delete button click
   */
  const handleDelete = () => {

    // Set the delete mode
    setDeleteMode(true);
    setEditMode(false);
    setNewMode(false);
  };

  /**
   * Handle the new button click
   */
  const handleNew = () => {

    // Set the new mode
    // disable edit and delete modes
    setSelectedIndex(-1);
    setNewMode(true);
    setEditMode(false);
    setDeleteMode(false);
  }

  /**
   * Handle the user selection
   * @param index The index of the selected user
   */
  const handleSelection = (index: number) => {

    // if we click on the same user, disable edit and delete modes
    if (index !== selectedIndex) {
      setEditMode(false);
      setDeleteMode(false);
      setNewMode(false);
    }

    // Set the selected index
    setSelectedIndex(index);
    console.log("Selected index: " + index);
  };


  /**
   * Fetch the users when the component mounts
   */
  useEffect(() => {
    console.log("Fetching users");
    // Fetch data using Axios when the component mounts
    axios.get('http://localhost:8080/api/getUsers') // Replace with your API endpoint
      .then(response => {

        // Set the state variable
        setUsers(response.data);
        console.log("Users fetched");
      })
      .catch(error => {

        // Handle error
        console.error('Error fetching data:', error);
      });
  }, []);

  // Return the JSX element
  return (
    <div className={styles.list}>
      <div className={styles.buttons}> 
        <button className={styles.new} id="new" onClick={handleNew}>New</button>
        {selectedIndex !== -1 && (
          <button className={styles.edit} id = "edit" onClick={handleEdit}>Edit</button>
        )}
        {selectedIndex !== -1 && (
          <button className={styles.delete} id="delete" onClick={handleDelete}>Delete</button>
        )
        }
      </div>
      
      {
        newMode && (
          <UserListItem
            key="new"
            user={{id: "-1", username: "", password: ""}}
            selected={true}
            editMode={false}
            deleteMode={false}
            newMode={true}
            setEditMode={setEditMode}
            setDeleteMode={setDeleteMode}
            setNewMode={setNewMode}
            onClick={() => handleSelection(-1)}
          />
        )
      }
      
      {users && users.map((user, index) => (
        <UserListItem
          key={user.id}
          user={user}
          selected={index === selectedIndex}
          editMode={editMode && index === selectedIndex}
          deleteMode={deleteMode && index === selectedIndex}
          newMode={false}
          setEditMode={setEditMode}
          setDeleteMode={setDeleteMode}
          setNewMode={setNewMode}
          onClick={() => handleSelection(index)}
        />
      ))}
    </div>
  );
};

export default UserList;
