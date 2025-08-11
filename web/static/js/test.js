// Test interface functionality

let questions = [];
let currentQuestionIndex = 0;
let answers = {};
let testStartTime = Date.now();
let questionStartTime = Date.now();
let testTimer = null;
let questionTimer = null;
let displayTimer = null;

async function initializeTest() {
    try {
        const response = await apiRequest('/api/questions?test_id=1');
        questions = response.data;
        
        if (questions.length === 0) {
            showNotification('No questions available', 'error');
            return;
        }
        
        setupQuestionNavigation();
        showQuestion(0);
        startTestTimer();
        
    } catch (error) {
        showNotification('Failed to load test questions', 'error');
        console.error('Error loading questions:', error);
    }
}

function setupQuestionNavigation() {
    const questionGrid = document.getElementById('questionGrid');
    questionGrid.innerHTML = '';
    
    for (let i = 0; i < questions.length; i++) {
        const navItem = document.createElement('div');
        navItem.className = 'question-nav-item';
        navItem.textContent = i + 1;
        navItem.onclick = () => goToQuestion(i);
        questionGrid.appendChild(navItem);
    }
}

function updateQuestionNavigation() {
    const navItems = document.querySelectorAll('.question-nav-item');
    navItems.forEach((item, index) => {
        item.classList.remove('current', 'answered');
        
        if (index === currentQuestionIndex) {
            item.classList.add('current');
        } else if (answers[index] !== undefined) {
            item.classList.add('answered');
        }
    });
}

function showQuestion(index) {
    if (index < 0 || index >= questions.length) return;
    
    currentQuestionIndex = index;
    questionStartTime = Date.now();
    
    const question = questions[index];
    
    // Update progress
    const progressFill = document.getElementById('progressFill');
    const progressText = document.getElementById('progressText');
    const progress = ((index + 1) / questions.length) * 100;
    
    progressFill.style.width = `${progress}%`;
    progressText.textContent = `Question ${index + 1} of ${questions.length}`;
    
    // Update question content
    document.getElementById('questionCategory').textContent = formatCategory(question.category);
    document.getElementById('questionText').textContent = question.question_text;
    
    // Clear previous question elements
    document.getElementById('questionDisplay').style.display = 'none';
    document.getElementById('questionOptions').innerHTML = '';
    document.getElementById('questionInput').style.display = 'none';
    
    // Setup question based on type
    setupQuestionByType(question);
    
    // Update navigation
    updateQuestionNavigation();
    updateNavigationButtons();
    
    // Start question timer if specified
    if (question.time_limit > 0) {
        startQuestionTimer(question.time_limit);
    }
    
    // Handle display time for memory questions
    if (question.display_time > 0) {
        showQuestionDisplay(question);
    }
}

function setupQuestionByType(question) {
    const optionsContainer = document.getElementById('questionOptions');
    const inputContainer = document.getElementById('questionInput');
    
    switch (question.question_type) {
        case 'multiple_choice':
            setupMultipleChoice(question, optionsContainer);
            break;
        case 'text_input':
            setupTextInput(question, inputContainer);
            break;
        case 'number_input':
            setupNumberInput(question, inputContainer);
            break;
        case 'key_sequence':
            setupKeySequence(question, inputContainer);
            break;
    }
}

function setupMultipleChoice(question, container) {
    const options = JSON.parse(question.options || '[]');
    
    options.forEach((option, index) => {
        const optionElement = document.createElement('div');
        optionElement.className = 'option';
        optionElement.innerHTML = `
            <input type="radio" name="answer" value="${String.fromCharCode(97 + index)}" id="option${index}">
            <label for="option${index}">${String.fromCharCode(65 + index)}) ${option}</label>
        `;
        
        optionElement.addEventListener('click', function() {
            const radio = this.querySelector('input[type="radio"]');
            radio.checked = true;
            saveAnswer(radio.value);
            
            // Update visual selection
            container.querySelectorAll('.option').forEach(opt => opt.classList.remove('selected'));
            this.classList.add('selected');
        });
        
        container.appendChild(optionElement);
    });
    
    // Restore previous answer
    const previousAnswer = answers[currentQuestionIndex];
    if (previousAnswer) {
        const radio = container.querySelector(`input[value="${previousAnswer}"]`);
        if (radio) {
            radio.checked = true;
            radio.closest('.option').classList.add('selected');
        }
    }
}

function setupTextInput(question, container) {
    container.innerHTML = `
        <input type="text" placeholder="Type your answer here..." id="textAnswer">
    `;
    container.style.display = 'block';
    
    const input = container.querySelector('input');
    input.addEventListener('input', function() {
        saveAnswer(this.value.trim());
    });
    
    // Restore previous answer
    const previousAnswer = answers[currentQuestionIndex];
    if (previousAnswer) {
        input.value = previousAnswer;
    }
    
    // Focus the input
    setTimeout(() => input.focus(), 100);
}

function setupNumberInput(question, container) {
    container.innerHTML = `
        <input type="number" placeholder="Enter number..." id="numberAnswer">
    `;
    container.style.display = 'block';
    
    const input = container.querySelector('input');
    input.addEventListener('input', function() {
        saveAnswer(this.value.trim());
    });
    
    // Restore previous answer
    const previousAnswer = answers[currentQuestionIndex];
    if (previousAnswer) {
        input.value = previousAnswer;
    }
    
    // Focus the input
    setTimeout(() => input.focus(), 100);
}

function setupKeySequence(question, container) {
    container.innerHTML = `
        <div class="key-sequence-input">
            <p>Press the key sequence as shown:</p>
            <div id="keySequenceDisplay"></div>
            <input type="text" readonly placeholder="Key sequence will appear here..." id="keySequenceAnswer">
        </div>
    `;
    container.style.display = 'block';
    
    const input = container.querySelector('#keySequenceAnswer');
    let keySequence = [];
    
    document.addEventListener('keydown', function(event) {
        if (currentQuestionIndex >= 0 && questions[currentQuestionIndex].question_type === 'key_sequence') {
            event.preventDefault();
            
            const key = event.key.toLowerCase();
            const mappedKey = mapKey(key);
            
            if (mappedKey) {
                keySequence.push(mappedKey);
                input.value = keySequence.join(',');
                saveAnswer(input.value);
            }
        }
    });
    
    // Restore previous answer
    const previousAnswer = answers[currentQuestionIndex];
    if (previousAnswer) {
        input.value = previousAnswer;
        keySequence = previousAnswer.split(',');
    }
}

function mapKey(key) {
    const keyMap = {
        'arrowup': 'up',
        'arrowdown': 'down',
        'arrowleft': 'left',
        'arrowright': 'right'
    };
    
    return keyMap[key] || key;
}

function showQuestionDisplay(question) {
    const displayContainer = document.getElementById('questionDisplay');
    displayContainer.style.display = 'block';
    
    // Show content based on question
    if (question.question_text.includes('7 2 9 4 6')) {
        displayContainer.innerHTML = '<div class="display-sequence">7 2 9 4 6</div>';
    } else if (question.question_text.includes('3 8 1 5')) {
        displayContainer.innerHTML = '<div class="display-sequence">3 8 1 5</div>';
    } else if (question.question_text.includes('L G K Q T')) {
        displayContainer.innerHTML = '<div class="display-sequence">L G K Q T</div>';
    }
    
    // Hide after display time
    displayTimer = setTimeout(() => {
        displayContainer.style.display = 'none';
    }, question.display_time * 1000);
}

function saveAnswer(answer) {
    answers[currentQuestionIndex] = answer;
    updateQuestionNavigation();
}

function startTestTimer() {
    const timerElement = document.getElementById('timer');
    
    testTimer = setInterval(() => {
        const elapsed = Math.floor((Date.now() - testStartTime) / 1000);
        timerElement.textContent = formatTime(elapsed);
    }, 1000);
}

function startQuestionTimer(timeLimit) {
    const timerElement = document.getElementById('questionTimer');
    let remaining = timeLimit;
    
    timerElement.textContent = `Time remaining: ${remaining}s`;
    
    questionTimer = setInterval(() => {
        remaining--;
        timerElement.textContent = `Time remaining: ${remaining}s`;
        
        if (remaining <= 0) {
            clearInterval(questionTimer);
            nextQuestion();
        }
    }, 1000);
}

function clearTimers() {
    if (questionTimer) {
        clearInterval(questionTimer);
        questionTimer = null;
    }
    
    if (displayTimer) {
        clearTimeout(displayTimer);
        displayTimer = null;
    }
    
    document.getElementById('questionTimer').textContent = '';
}

function updateNavigationButtons() {
    const prevBtn = document.getElementById('prevBtn');
    const nextBtn = document.getElementById('nextBtn');
    const submitBtn = document.getElementById('submitBtn');
    
    prevBtn.disabled = currentQuestionIndex === 0;
    
    if (currentQuestionIndex === questions.length - 1) {
        nextBtn.style.display = 'none';
        submitBtn.style.display = 'inline-flex';
    } else {
        nextBtn.style.display = 'inline-flex';
        submitBtn.style.display = 'none';
    }
}

function previousQuestion() {
    if (currentQuestionIndex > 0) {
        clearTimers();
        showQuestion(currentQuestionIndex - 1);
    }
}

function nextQuestion() {
    if (currentQuestionIndex < questions.length - 1) {
        clearTimers();
        showQuestion(currentQuestionIndex + 1);
    }
}

function goToQuestion(index) {
    if (index >= 0 && index < questions.length) {
        clearTimers();
        showQuestion(index);
    }
}

function submitTest() {
    const modal = document.getElementById('submitModal');
    modal.style.display = 'block';
}

function closeSubmitModal() {
    const modal = document.getElementById('submitModal');
    modal.style.display = 'none';
}

async function confirmSubmit() {
    closeSubmitModal();
    
    const submitButton = document.getElementById('submitBtn');
    showLoading(submitButton);
    
    try {
        const testAnswers = [];
        
        for (let i = 0; i < questions.length; i++) {
            const answer = answers[i] || '';
            const responseTime = 5000; // Default response time, would be tracked in real implementation
            
            testAnswers.push({
                question_id: questions[i].id,
                user_answer: answer,
                response_time: responseTime
            });
        }
        
        const totalTime = Math.floor((Date.now() - testStartTime) / 1000);
        
        const submitData = {
            test_id: 1,
            answers: testAnswers,
            time_taken: totalTime
        };
        
        const response = await apiRequest('/api/submit', {
            method: 'POST',
            body: JSON.stringify(submitData)
        });
        
        if (response.success) {
            showNotification('Test submitted successfully!', 'success');
            
            // Clear test timer
            if (testTimer) {
                clearInterval(testTimer);
            }
            
            // Redirect to results
            setTimeout(() => {
                window.location.href = '/results';
            }, 2000);
        }
        
    } catch (error) {
        showNotification('Failed to submit test', 'error');
        console.error('Submit error:', error);
    } finally {
        hideLoading(submitButton);
    }
}

function formatCategory(category) {
    return category.split('_').map(word => 
        word.charAt(0).toUpperCase() + word.slice(1)
    ).join(' ');
}

// Prevent accidental page navigation
window.addEventListener('beforeunload', function(event) {
    if (questions.length > 0 && Object.keys(answers).length > 0) {
        event.preventDefault();
        event.returnValue = 'You have unsaved test progress. Are you sure you want to leave?';
        return event.returnValue;
    }
});

// Close modal when clicking outside
window.addEventListener('click', function(event) {
    const modal = document.getElementById('submitModal');
    if (event.target === modal) {
        closeSubmitModal();
    }
});