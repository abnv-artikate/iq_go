# Cognitive Assessment Test Application

A comprehensive web application built with Go (Gin framework) and PostgreSQL for conducting cognitive assessments. The application evaluates users across five key cognitive domains: Analytical Reasoning, Working Memory, Processing Speed, Attention & Focus, and Emotional Regulation.

## Features

- **User Authentication**: JWT-based authentication with secure login/registration
- **Comprehensive Testing**: 50 questions across 5 cognitive domains
- **Multiple Question Types**: Multiple choice, text input, number input, and key sequence
- **Real-time Testing**: Timed questions with progress tracking
- **Results Dashboard**: Detailed performance analytics and history
- **Responsive Design**: Mobile-friendly interface
- **Docker Support**: Easy deployment with Docker Compose

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Docker & Docker Compose (optional)

### Using Docker Compose

```bash
# Clone the repository
git clone <repository-url>
cd iq-go

# Start the application
docker-compose up -d

# The application will be available at http://localhost:8080
```

### Manual Setup

1. **Setup Database**
```bash
# Create PostgreSQL database
createdb cognitiondb
```

2. **Configure Environment**
```bash
# Copy and edit environment variables
cp .env.example .env
# Edit DATABASE_URL, JWT_SECRET, and PORT in .env file
```

3. **Install Dependencies**
```bash
go mod download
```

4. **Seed Questions**
```bash
go run scripts/seed_questions.go
```

5. **Start Application**
```bash
go run cmd/server/main.go
```

## API Endpoints

### Authentication
- `POST /api/register` - User registration
- `POST /api/login` - User login
- `POST /api/logout` - User logout

### Test Management
- `GET /api/questions` - Get test questions
- `POST /api/submit` - Submit test answers

### Results
- `GET /api/results` - Get user's test results
- `GET /api/results/:id` - Get specific test result details

## Project Structure

```
iq-go/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── auth/           # Authentication middleware
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and migrations
│   ├── handlers/       # HTTP request handlers
│   ├── models/         # Data models
│   ├── services/       # Business logic
│   └── utils/          # Utility functions
├── web/                # Frontend assets
│   ├── static/         # CSS, JS, images
│   └── templates/      # HTML templates
├── scripts/            # Database seeding scripts
└── docker-compose.yml  # Docker configuration
```

## Database Schema

### Users
- ID, Email, Password (hashed)
- First Name, Last Name
- Created/Updated timestamps

### Tests
- ID, Name, Description, Duration
- Created/Updated timestamps

### Questions
- ID, Test ID, Question Text, Type, Category
- Options (JSON), Correct Answer, Time Limits
- Order Index, Display Time

### Test Results
- ID, User ID, Test ID, Score, Total Questions
- Time Taken, Start/Completion timestamps

### Answers
- ID, Test Result ID, Question ID
- User Answer, Correctness, Response Time

## Question Types

1. **Multiple Choice**: Standard options (A, B, C, D)
2. **Text Input**: Free-form text responses
3. **Number Input**: Numeric answers
4. **Key Sequence**: Keyboard input sequences

## Cognitive Domains

1. **Analytical Reasoning** (Questions 1-10)
   - Logic chains, pattern recognition, analogies
   - Problem-solving and deductive reasoning

2. **Working Memory** (Questions 11-20)
   - Digit sequences, memory recall tasks
   - Information retention and manipulation

3. **Processing Speed** (Questions 21-30)
   - Quick calculations, spelling, word recognition
   - Rapid task completion under time pressure

4. **Attention & Focus** (Questions 31-40)
   - Pattern detection, counting tasks
   - Sustained concentration abilities

5. **Emotional Regulation** (Questions 41-50)
   - Workplace scenarios, stress management
   - Emotional intelligence and coping strategies

## Development

### Adding New Questions
1. Update `scripts/seed_questions.go`
2. Define question with proper type and category
3. Run seeding script to update database

### Extending Question Types
1. Add new type to `models/question.go`
2. Update evaluation logic in `services/test.go`
3. Implement frontend handling in `test.js`

### Environment Variables
```env
DATABASE_URL=postgres://username:password@localhost:5432/cognitiondb?sslmode=disable
JWT_SECRET=your-secure-secret-key
PORT=8080
```

## Testing

Run the application and navigate to:
- Registration: `http://localhost:8080/register`
- Login: `http://localhost:8080/login`
- Dashboard: `http://localhost:8080/`
- Take Test: `http://localhost:8080/test`
- View Results: `http://localhost:8080/results`

## Production Deployment

1. **Security**: Update JWT secret and database credentials
2. **HTTPS**: Configure SSL/TLS certificates
3. **Database**: Use managed PostgreSQL service
4. **Monitoring**: Add logging and health checks
5. **Scaling**: Configure load balancing if needed

## License

This project is licensed under the MIT License.