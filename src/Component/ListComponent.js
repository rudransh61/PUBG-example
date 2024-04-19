// Component/ListComponent.js

// Define ListComponent
export function ListComponent(todos, removeTodo) {
    // Create a list element
    const list = {
        type: "ul",
        content: todos.map((todo, index) => ({
            type: "li",
            content: [
                { type: "span", content: todo },
                {
                    type: "button",
                    content: "Remove",
                    onclick: () => removeTodo(index),
                },
            ],
        })),
    };

    return list;
}
