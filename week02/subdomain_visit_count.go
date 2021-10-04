package week02

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/subdomain-visit-count/
// leetcode 811 子域名计数访问

type domainTree struct {
	domain   string
	count    int
	children map[string]*domainTree
}

//树实现
func subdomainVisits2(cpdomains []string) []string {
	var root = &domainTree{children: make(map[string]*domainTree)}
	for _, domain := range cpdomains {
		var domains = strings.Split(domain, " ")
		var num, _ = strconv.Atoi(domains[0])
		domains = strings.Split(domains[1], ".")
		var node = root
		for i := len(domains) - 1; i >= 0; i-- {
			if child, ok := node.children[domains[i]]; ok {
				child.count += num
				node = child
			} else {
				if node.children == nil {
					node.children = make(map[string]*domainTree)
				}
				var child = &domainTree{
					domain: domains[i],
					count:  num,
				}
				node.children[domains[i]] = child
				node = child
			}
		}
	}
	return getTree(root, "")
}

func getTree(node *domainTree, prefix string) []string {
	var ant []string
	if len(prefix) == 0 {
		prefix = node.domain
	} else {
		prefix = node.domain + "." + prefix
	}
	if node.count > 0 {
		ant = []string{strconv.Itoa(node.count) + " " + prefix}
	}
	if len(node.children) > 0 {
		for _, child := range node.children {
			ant = append(ant, getTree(child, prefix)...)
		}
	}
	return ant
}

// ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
// 哈希表实现计数
func subdomainVisits(cpdomains []string) []string {
	var ant = make(map[string]int, len(cpdomains))
	for _, domain := range cpdomains {
		var domains = strings.Split(domain, " ")
		var num, _ = strconv.Atoi(domains[0])
		domains = strings.Split(domains[1], ".")
		var domainStr string
		for i := len(domains) - 1; i >= 0; i-- {
			if len(domainStr) == 0 {
				domainStr += domains[i]
			} else {
				domainStr = domains[i] + "." + domainStr
			}

			ant[domainStr] += num
		}
	}
	var antStr []string
	for str, num := range ant {
		antStr = append(antStr, strconv.Itoa(num)+" "+str)
	}
	return antStr
}
