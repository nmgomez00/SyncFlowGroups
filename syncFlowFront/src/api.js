import axios from "axios";
const API_BASE = import.meta.env.VITE_API_BASE || "http://localhost:8080";
export function createUser({ name, email, profilePhotoURL }) {
  return axios
    .post(`${API_BASE}/users`, { name, email, profilePhotoURL })
    .then((res) => res.data);
}
// User helpers: create a local user and store id in localStorage
export async function ensureLocalUser() {
  let uid = localStorage.getItem("syncflow:userID");
  if (uid) {
    // Check if user exists in backend
    try {
      const users = await axios
        .get(`${API_BASE}/users`)
        .then((res) => res.data);
      if (users.some((u) => u.id === uid || u.userID === uid)) {
        return uid;
      } else {
        // User not found, remove from localStorage
        localStorage.removeItem("syncflow:userID");
      }
    } catch (err) {
      console.warn("[ensureLocalUser] Could not verify user in backend", err);

      // fallback: return uid
      //return uid;
    }
  }
  const random = Math.random().toString(36).slice(2, 9);
  const payload = {
    name: `Local User ${random}`,
    email: `${random}@local.dev`,
    profilePhotoURL: "",
  };
  console.log("[ensureLocalUser] Sending POST /users", payload);
  const { data } = await axios.post(`${API_BASE}/users`, payload);
  console.log("[ensureLocalUser] Received response:", data);
  uid = data.userID;
  localStorage.setItem("syncflow:userID", uid);
  console.log("[ensureLocalUser] Stored userID:", uid);
  return uid;
}

export function getGroups() {
  return axios.get(`${API_BASE}/groups`).then((res) => res.data);
}
export function createGroup({
  name,
  description,
  privacy = "PUBLIC",
  state = "ACTIVE",
  userID,
}) {
  return axios
    .post(`${API_BASE}/groups`, { name, description, privacy, state, userID })
    .then((res) => ({
      id: res.data.groupID,
      name,
      description,
      privacy,
      state,
      userID,
    }));
}
export function deleteGroup(groupID) {
  return axios.delete(`${API_BASE}/groups/${groupID}`);
}
export function joinGroup(
  groupID,
  { userID, role = "USER", state = "JOINED" }
) {
  return axios.post(`${API_BASE}/groups/${groupID}/join`, {
    userID,
    role,
    state,
  });
}
export function getCategories(groupID) {
  return axios
    .get(`${API_BASE}/groups/${groupID}/categories`)
    .then((res) => res.data);
}
export function createCategory(groupID, { name, description, userCreatedID }) {
  return axios
    .post(`${API_BASE}/groups/${groupID}/categories`, {
      name,
      description,
      userCreatedID,
    })
    .then((res) => res.data);
}
export function deleteCategory(groupID, categoryID) {
  return axios.delete(`${API_BASE}/groups/${groupID}/categories/${categoryID}`);
}
export function getChannels(groupID) {
  return axios
    .get(`${API_BASE}/groups/${groupID}/channels`)
    .then((res) => res.data);
}
export function createChannel(
  groupID,
  categoryID,
  { name, description, channelState = "ACTIVE", userID }
) {
  return axios
    .post(`${API_BASE}/groups/${groupID}/categories/${categoryID}/channels`, {
      name,
      description,
      channelState,
      userID,
    })
    .then((res) => res.data);
}
export function deleteChannel(groupID, categoryID, channelID) {
  return axios.delete(
    `${API_BASE}/groups/${groupID}/categories/${categoryID}/channels/${channelID}`
  );
}
export function getChannelsByCategory(groupID, categoryID) {
  return axios
    .get(`${API_BASE}/groups/${groupID}/categories/${categoryID}/channels`)
    .then((res) => res.data);
}
export function getAllUsersByGroup(groupID) {
  return axios
    .get(`${API_BASE}/groups/${groupID}/users`)
    .then((res) => res.data);
}
export function getUsers() {
  return axios.get(`${API_BASE}/users`).then((res) => res.data);
}
export function deleteUser(userID) {
  return axios.delete(`${API_BASE}/users/${userID}`);
}
export function leaveGroup(groupID, userID) {
  return axios.delete(`${API_BASE}/groups/${groupID}/users/${userID}`);
}
export function changeUserRole(groupID, userID, role) {
  return axios.patch(`${API_BASE}/groups/${groupID}/users/${userID}/role`, {
    role,
  });
}
