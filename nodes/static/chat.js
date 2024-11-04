document.addEventListener('DOMContentLoaded', () => {
    const chatContainer = document.getElementById('chat-container');
    const messageForm = document.getElementById('message-form');
    const messageInput = document.getElementById('message-input');
    let ws = null;

    function connectWebSocket() {
        ws = new WebSocket('ws://localhost:6969/ws');

        ws.onopen = () => {
            console.log('Connected to WebSocket');
            addSystemMessage('Connected to chat');
        };

        ws.onmessage = (event) => {
            const message = JSON.parse(event.data);
            addMessage(message, false);
        };

        ws.onclose = () => {
            console.log('Disconnected from WebSocket');
            addSystemMessage('Disconnected from chat. Attempting to reconnect...');
            setTimeout(connectWebSocket, 5000);
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
            addSystemMessage('Error connecting to chat');
        };
    }

    function addMessage(message, isSent) {
        const messageDiv = document.createElement('div');
        messageDiv.className = `message ${isSent ? 'sent' : 'received'}`;

        const username = document.createElement('div');
        username.className = 'username';
        username.textContent = message.username || 'Anonymous';

        const content = document.createElement('div');
        content.className = 'content';
        content.textContent = message.content;

        const timestamp = document.createElement('div');
        timestamp.className = 'timestamp';

        messageDiv.appendChild(username);
        messageDiv.appendChild(content);
        messageDiv.appendChild(timestamp);

        chatContainer.appendChild(messageDiv);
        chatContainer.scrollTop = chatContainer.scrollHeight;
    }

    function addSystemMessage(text) {
        const messageDiv = document.createElement('div');
        messageDiv.className = 'message';
        messageDiv.style.backgroundColor = 'var(--secondary-color)';
        messageDiv.style.color = 'var(--highlight-color)';
        messageDiv.style.textAlign = 'center';
        messageDiv.style.alignSelf = 'center';
        messageDiv.textContent = text;
        chatContainer.appendChild(messageDiv);
        chatContainer.scrollTop = chatContainer.scrollHeight;
    }

    messageForm.addEventListener('submit', (e) => {
        e.preventDefault();

        const content = messageInput.value.trim();
        if (!content) return;

        const message = {
            content,
            username: 'You', // This could be replaced with actual username
            timestamp: Date.now()
        };

        if (ws && ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify(message));
        } else {
            addSystemMessage('Unable to send message: Not connected');
        }

        messageInput.value = '';
    });

    // Initialize WebSocket connection
    connectWebSocket();
});