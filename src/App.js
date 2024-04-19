// App.js
import { createState } from "./state.js";
import { ListComponent } from "./Component/ListComponent.js";

// Create reactive state
const initialState = {
    count: 0,
    todos: [], // Initialize todos array
};
const state = createState(initialState);

// Define components
const Increment = () => {
    state.setState({ count: state.getState().count + 1 });
    console.log(state.getState().count);
};

const Decrement = () => {
    state.setState({ count: state.getState().count - 1 });
    console.log(state.getState().count);
};

const SetTodo = async () => {
    const todoText = prompt("Enter a new todo:");
    if (todoText) {
        try {
            console.log(todoText)
            const response = await fetch("http://localhost:6969/set", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ key: todoText, value:"cfhnjcn", ttl:"1000s" }),
            });
            const data = await response.json();
            const updatedTodos = [...state.getState().todos, data];
            state.setState({ todos: updatedTodos });
        } catch (error) {
            console.error("Error setting todo:", error);
        }
    }
};

const UpdateTodos = async () => {
    try {
        const response = await fetch("http://localhost:6969/list");
        const data = await response.json();
        console.log(data)
        state.setState({ todos: data.keys });
    } catch (error) {
        console.error("Error updating todos:", error);
    }
};

// Define your main App component
export function App() {
    // UpdateTodos()
    return {
        type: "div",
        content: [
            {
                type: "p",
                id: "text",
                content: `Count: ${state.getState().count}`,
                style: {
                    color: "red",
                    fontSize: `20px`,
                    backgroundColor: "lightblue",
                },
            },
            {
                type: "button",
                content: "Increment",
                onclick: Increment,
            },
            {
                type: "button",
                content: "Decrement",
                onclick: Decrement,
            },
            {
                type: "button",
                content: "Set Todo",
                onclick: SetTodo,
            },
            {
                type: "button",
                content: "Update Todos",
                onclick: UpdateTodos,
            },
            ListComponent(state.getState().todos) // Pass todos to ListComponent
        ],
    };
}
