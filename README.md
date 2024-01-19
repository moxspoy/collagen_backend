# collagen_backend
a golang project powered by gin as web service for collagen

# Collagen: Student Social Media Platform

<img src="https://iconape.com/wp-content/png_logo_vector/gopher.png" alt="drawing" width="200"/>

Collagen is a social media platform tailored specifically for students, designed to support their academic journey while providing a space for collaboration, discussion, and sharing experiences. Think of it as the "Reddit for students," built with the efficiency and power of the Gin framework in Golang.

## Features

- **Communities**: Join communities related to your courses, interests, or campus events.
- **Posts and Discussions**: Share your thoughts, ask questions, and engage in discussions with fellow students.
- **Resource Sharing**: Easily share and discover study materials, lecture notes, and helpful resources.
- **Events**: Stay updated on campus events, club meetings, and other activities.
- **Connect with Peers**: Build your network and connect with classmates to enhance your learning experience.
- **Anonymous Mode**: Discuss sensitive topics or ask questions anonymously to encourage open communication.

## Getting Started

To run Collagen locally, follow these simple steps:

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/moxspoy/collagen_backend.git
   cd collagen_backend
   ```
[config](.git%2Fconfig)
2. **Install Dependencies:**
   ```bash
   go mod download
   ```

3. **Run the Application:**[config](.git%2Fconfig)
   ```bash
   go run main.go
   ```

   Collagen will be accessible at `http://localhost:8083`.

## Contributing

We welcome contributions to make Collagen even better! To contribute, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Submit a pull request.

## Tech Stack

- **Gin Framework**: Fast and lightweight Golang web framework.
- **Mysql**: MySQL database for storing user data and posts.
- **Gow**: File watcher
- **SQLite**: Database testing

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Hat tip to the contributors who make Collagen possible.
- Inspired by the desire to create a supportive community for students.

Feel free to explore, contribute, and make Collagen a vibrant space for students worldwide!
