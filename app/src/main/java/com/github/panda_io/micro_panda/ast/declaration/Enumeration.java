package com.github.panda_io.micro_panda.ast.declaration;

import com.github.panda_io.micro_panda.ast.type.Type;
import com.github.panda_io.micro_panda.ast.Context;

public class Enumeration extends Declaration {
    public boolean isConstant()  {
        return false;
    }

    public Type getType()  {
        return null;
    }

    public void resolveType(Context context) {
    }

	public void validate(Context context) {
	}
	
	public boolean hasMember(String memberName) {
		return false;
	}
	/*
type Enum struct {
	DeclarationBase
	Members []*Variable
	Values  []uint8
}

func (e *Enum) AddMember(m *Variable) error {
	if e.HasMember(m.Name.Name) {
		return fmt.Errorf("%s redeclared", m.Name.Name)
	}
	e.Members = append(e.Members, m)
	return nil
}

func (e *Enum) HasMember(name string) bool {
	for _, v := range e.Members {
		if v.Name.Name == name {
			return true
		}
	}
	return false
}

func (e *Enum) ResolveType(c ast.Context) {
}

func (e *Enum) Validate(c ast.Context) {
	var index int
	for _, v := range e.Members {
		if index >= 256 {
			c.Error(v.GetPosition(), "enum value shoud be less than 256")
		}
		if v.Value == nil {
			e.Values = append(e.Values, uint8(index))
			index++
		} else {
			if literal, ok := v.Value.(*expression.Literal); ok && literal.Token == token.INT {
				if i, _ := strconv.Atoi(literal.Value); i >= index {
					index = i
					e.Values = append(e.Values, uint8(index))
					index++
				} else {
					c.Error(v.GetPosition(), fmt.Sprintf("enum value here should be greater than %d.", i-1))
				}
			} else {
				c.Error(v.GetPosition(), "enum value must be const integer.")
			}
		}
	}
}*/

}
