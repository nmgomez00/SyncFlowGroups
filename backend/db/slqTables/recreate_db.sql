DROP TABLE IF EXISTS "UserGroupRequest", "UserGroup", "Channel", "Category", "Group", "User";


-- 1. Enum para el estado de la membresía en un grupo (UserGroup)
CREATE TYPE GroupMembershipStates AS ENUM (
    'PENDING',
    'REJECTED',
    'CANCELLED',
    'JOINED'
);

-- 2. Enum para el rol dentro del grupo (UserGroup)
CREATE TYPE GroupRoles AS ENUM (
    'OWNER',
    'ADMIN',
    'USER'
);

-- 3. Enum para la visibilidad del grupo (Group)
CREATE TYPE GroupPrivacy AS ENUM (
    'PUBLIC',
    'PRIVATE'
);

-- 4. Enum para el estado de un grupo (Group)
CREATE TYPE GroupStates AS ENUM (
    'ACTIVE',
    'DELETED',
    'BLOCKED',
    'INACTIVE'
);

-- 5. Enum para el estado de un canal (Channel) - Aunque no se ve directamente en Channel, se infiere del contexto.
CREATE TYPE ChannelStates AS ENUM (
    'ACTIVE',
    'BANNED',
    'LEFT'
);

CREATE TABLE "User" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    is_blocked BOOLEAN DEFAULT FALSE,
    profile_photo_url TEXT
);

CREATE TABLE "Group" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    user_created_id UUID NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_activity_date TIMESTAMP WITH TIME ZONE,
    -- Usa los ENUMs definidos
    privacy GroupPrivacy NOT NULL, 
    state GroupStates NOT NULL DEFAULT 'ACTIVE',
    FOREIGN KEY (user_created_id) REFERENCES "User"(id) ON DELETE CASCADE
);
CREATE TABLE "Category" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    group_id UUID NOT NULL,
    user_created_id UUID NOT NULL,
    FOREIGN KEY (group_id) REFERENCES "Group"(id) ON DELETE CASCADE,
    FOREIGN KEY (user_created_id) REFERENCES "User"(id) ON DELETE CASCADE
);
CREATE TABLE "Channel" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    group_id UUID NOT NULL,
    category_id UUID NOT NULL,
    -- Usa el ENUM definido
    channel_state ChannelStates NOT NULL DEFAULT 'ACTIVE', 
    
    FOREIGN KEY (group_id) REFERENCES "Group"(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES "Category"(id) ON DELETE CASCADE
);
CREATE TABLE "UserGroup" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    group_id UUID NOT NULL,
    -- Usa los ENUMs definidos
    role GroupRoles NOT NULL, 
    state GroupMembershipStates NOT NULL DEFAULT 'JOINED',
    
    -- Restricción para asegurar que un usuario solo tenga una relación con un grupo
    UNIQUE (user_id, group_id),
    
    FOREIGN KEY (user_id) REFERENCES "User"(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES "Group"(id) ON DELETE CASCADE
);
CREATE TABLE "UserGroupRequest" (
    id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    status GroupMembershipStates NOT NULL DEFAULT 'PENDING', -- Reutilizamos el ENUM para el estado de la solicitud
    request_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Restricción para evitar que un usuario solicite unirse al mismo grupo dos veces
    UNIQUE (group_id, user_id),
    
    FOREIGN KEY (group_id) REFERENCES "Group"(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES "User"(id) ON DELETE CASCADE
);
