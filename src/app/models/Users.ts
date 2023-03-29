

export interface Chat {
  id: string;
  lastMessage?: string;
  lastMessageDate?: Date;
  userIds: string[];
  users: ProfileUser[];

  // Not stored, only for display
  chatPic?: string;
  chatName?: string;
}

export interface Message {
  text: string;
  senderId: string;
  sentDate: Date;
}

export interface ProfileUser {
  uid: string;
  email?: string;
  displayName?: string;
  photoURL?: string;
  firstName?: string;
  lastName?: string;
  phone?: string;
  address?: string;
}
