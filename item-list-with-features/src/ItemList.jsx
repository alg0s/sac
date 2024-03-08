import React, { useState } from 'react';
import './ItemList.css';

export const ItemList = ({ data = [] }) => {
  const [items, setItems] = useState(data);
  const [input, setInput] = useState('');
  const [editIndex, setEditIndex] = useState(null);

  const handleAdd = (event) => {
    event.preventDefault();
    if (editIndex !== null) {
      setItems(items.map((item, index) => index === editIndex ? input : item));
      setEditIndex(null);
    } else {
      setItems([...items, input]);
    }
    setInput('');
  };

  const handleDelete = (index) => {
    if (window.confirm('Are you sure you want to delete this item?')) {
      setItems(items.filter((item, i) => i !== index));
    }
  };

  const handleEdit = (index) => {
    setInput(items[index]);
    setEditIndex(index);
  };

  return (
    <div>
      {items.map((item, index) => (
        <div className="columns container">
          <p className="column item">{item}</p>
          <div className="buttons column hide">
            <button onClick={() => handleDelete(index)} className="button is-small is-danger">Delete</button>
            <button onClick={() => handleEdit(index)} className="button is-small is-info">Edit</button>
          </div>
        </div>
      ))}
      <form onSubmit={handleAdd}>
        <input className="input" type="text" value={input} onChange={e => setInput(e.target.value)} />
        <button className="button is-primary my-2" type="submit">Submit</button>
      </form>
    </div>
  );
};