import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styles from './list.module.css';

interface User {
  id: string;
  username: string;
  password: string;
}

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

const UserListItem: React.FC<UserListItemProps> = (props) => {
  const { user, selected, editMode, deleteMode, newMode, setEditMode, setDeleteMode, setNewMode, onClick } = props;
  const [id, setId] = useState<string>(user.id);
  const [username, setUsername] = useState<string>(user.username);
  const [password, setPassword] = useState<string>(user.password);
  
  const deleteUser = async () => {
    console.log("Deleting user");
    await axios.delete(`http://localhost:8080/deleteUser?id=${id}`, 
    ).then(response => {
      console.log(response);
      alert(response)
    }).catch(error => {
      console.error('Error deleting user:', error);
    });

    window.location.reload();
  }

  const updateUser = async () => {
    console.log("Updating user");
    await axios.put(`http://localhost:8080/updateUser?id=${id}&username=${username}&password=${password}`, 
    ).then(response => {
      console.log(response.data.message);
    }).catch(error => {
      console.error('Error updating user:', error);
    });

    window.location.reload();
  }

  const createNewUser = async () => {
    console.log("Creating new user");
    await axios.post(`http://localhost:8080/saveUser?username=${username}&password=${password}`, 
    ).then(response => {
      console.log(response.data.message);
    }).catch(error => {
      console.error('Error creating new user:', error);
    });

    window.location.reload();
  }
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

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [editMode, setEditMode] = useState<boolean>(false);
  const [deleteMode, setDeleteMode] = useState<boolean>(false);
  const [newMode, setNewMode] = useState<boolean>(false);
  const [selectedIndex, setSelectedIndex] = useState<number>(-1);

  const handleEdit = () => {
    setEditMode(true);
    setDeleteMode(false);
    setNewMode(false);
  };

  const handleDelete = () => {
    setDeleteMode(true);
    setEditMode(false);
    setNewMode(false);
  };

  const handleNew = () => {
    setSelectedIndex(-1);
    setNewMode(true);
    setEditMode(false);
    setDeleteMode(false);
  }

  const handleSelection = (index: number) => {

    if (index !== selectedIndex) {
      setEditMode(false);
      setDeleteMode(false);
    }
    setSelectedIndex(index);
    console.log("Selected index: " + index);
  };

  useEffect(() => {
    console.log("Fetching users");
    // Fetch data using Axios when the component mounts
    axios.get('http://localhost:8080/getUsers') // Replace with your API endpoint
      .then(response => {
        setUsers(response.data);
        console.log("Users fetched");
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <div className={styles.list}>
      <div className={styles.buttons}> 
        <button onClick={handleNew}>New</button>
        {selectedIndex !== -1 && (
          <button onClick={handleEdit}>Edit</button>
        )}
        {selectedIndex !== -1 && (
          <button onClick={handleDelete}>Delete</button>
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
