package utils

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func UpdateYAMLFile(filename string, updates map[string]interface{}, createIfAbsent bool) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var root yaml.Node
	if err := yaml.Unmarshal(content, &root); err != nil {
		return fmt.Errorf("failed to parse yaml: %w", err)
	}

	println("root is:", root.Content)
	println("root is:", root.Kind)

	for path, value := range updates {
		keys := strings.Split(path, ".")
		if err := updateYAMLNode(&root, keys, value, createIfAbsent); err != nil {
			return fmt.Errorf("failed to update path '%s': %w", path, err)
		}
	}

	out, err := yaml.Marshal(&root)
	if err != nil {
		return fmt.Errorf("failed to marshal yaml: %w", err)
	}

	return os.WriteFile(filename, out, 0644)
}

func updateYAML(node *yaml.Node, path []string, newValue interface{}) {
	if len(path) == 0 || node == nil {
		return
	}

	if node.Kind == yaml.DocumentNode {
		for _, child := range node.Content {
			updateYAML(child, path, newValue)
		}
		return
	}

	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			k := node.Content[i]
			v := node.Content[i+1]

			if k.Value == path[0] {
				if len(path) == 1 {
					v.Kind = yaml.ScalarNode
					v.Value = fmt.Sprintf("%v", newValue)
					return
				}
				updateYAML(v, path[1:], newValue)
			}
		}
	}
}

func updateYAMLNode(root *yaml.Node, parts []string, value interface{}, createIfAbsent bool) error {
	curr := root.Content[0]

	for i, key := range parts {
		if curr.Kind != yaml.MappingNode {
			return fmt.Errorf("path '%s' is not a mapping node", strings.Join(parts[:i], "."))
		}

		found := false
		for j := 0; j < len(curr.Content); j += 2 {
			k := curr.Content[j]
			v := curr.Content[j+1]
			if k.Value == key {
				if i == len(parts)-1 {
					v.Value = fmt.Sprintf("%v", value)
					return nil
				}
				curr = v
				found = true
				break
			}
		}

		if !found {
			if !createIfAbsent {
				return fmt.Errorf("path not found: %s", strings.Join(parts[:i+1], "."))
			}
			newKey := &yaml.Node{Kind: yaml.ScalarNode, Value: key}
			newVal := &yaml.Node{Kind: yaml.MappingNode}
			curr.Content = append(curr.Content, newKey, newVal)
			curr = newVal

			if i == len(parts)-1 {
				curr.Kind = yaml.ScalarNode
				curr.Value = fmt.Sprintf("%v", value)
				return nil
			}
		}
	}

	return nil
}
