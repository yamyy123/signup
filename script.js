document.addEventListener('DOMContentLoaded', function () {
    const signupForm = document.getElementById('signup-form');
    const signupButton = document.getElementById('signup-button');
  
    signupButton.addEventListener('click', function () {
      const username = document.getElementById('username').value;
      const email = document.getElementById('email').value;
      const password = document.getElementById('password').value;
  
      // Send the data to the backend for registration
      fetch('/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, email, password }),
      })
        .then(response => response.json())
        .then(data => {
          // Handle the response from the server
          if (data.success) {
            alert('Registration successful!');
          } else {
            alert('Registration failed. Please try again.');
          }
        });
    });
  });
  